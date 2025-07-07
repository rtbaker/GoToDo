<?php

declare(strict_types=1);

namespace App\Controller;

use Exception;
use App\Entity\ToDo;
use App\Entity\User;
use OpenApi\Attributes as OA;
use Nelmio\ApiDocBundle\Attribute\Model;
use App\Repository\ToDoRepository;
use Doctrine\ORM\EntityManagerInterface;
use Doctrine\DBAL\Exception\DriverException;
use Symfony\Component\HttpFoundation\Request;
use Symfony\Component\HttpFoundation\Response;
use Symfony\Component\Routing\Attribute\Route;
use Symfony\Component\Serializer\SerializerInterface;
use Symfony\Component\HttpKernel\Exception\HttpException;
use Symfony\Component\Security\Http\Attribute\CurrentUser;
use Symfony\Component\Validator\Validator\ValidatorInterface;
use Symfony\Component\Serializer\Normalizer\AbstractNormalizer;
use Symfony\Bundle\FrameworkBundle\Controller\AbstractController;

 #[Route('/api/1.0/todos', name: 'app_todo_')]
final class TodoController extends AbstractController
{
    use ErrorTrait;

    public function __construct(
        private ToDoRepository $toDoRepository,
        private SerializerInterface $serializer,
        private EntityManagerInterface $entityManager,
        private ValidatorInterface $validator
    ) {
    }

    #[Route('', name: 'read', methods: ['GET'])]
    #[OA\Response(
        response: 200,
        description: 'Todo created successfully',
        content: new OA\JsonContent(
            type: 'array',
            items: new OA\Items(new Model(type: ToDo::class, groups: ['read']))
        )
    )]
    public function read(#[CurrentUser] User $user): Response
    {
        $todos = $this->toDoRepository->findByUser($user);

        return $this->json(
            $todos,
            Response::HTTP_OK,
            [],
            ['groups' => ['read']]
        );
    }

    #[Route('', name: 'create', methods: ['POST'])]
    #[OA\RequestBody(
        description: 'JSON request body.',
        required: true,
        content: new OA\JsonContent(
            ref: new Model(type: ToDo::class, groups: ['create'])
        )
    )]
    #[OA\Response(
        response: 200,
        description: 'Todo created successfully',
        content: new OA\JsonContent(ref: new Model(type: ToDo::class, groups: ['read']))
    )]
    public function create(
        #[CurrentUser] User $user,
        Request $request
    ): Response {
        try {
            $todo = $this->serializer->deserialize(
                $request->getContent(),
                ToDo::class,
                'json',
                ['allow_extra_attributes' => false, 'groups' => ['create']]
            );

            $todo->setUser($user);

            $errors = $this->validator->validate($todo);
            if (count($errors)) {
                $this->throwValidatorException($errors);
            }

            $this->toDoRepository->add($todo, true);
        } catch (DriverException $e) {
            return $this->returnJsonError($e->getMessage(), Response::HTTP_INTERNAL_SERVER_ERROR);
        } catch (Exception $e) {
            return $this->returnJsonError(
                $e->getMessage(),
                $e->getCode() !== 0 ? $e->getCode() : Response::HTTP_UNPROCESSABLE_ENTITY
            );
        }


        return $this->json(
            $todo,
            Response::HTTP_OK,
            [],
            ['groups' => ['read']]
        );
    }

    #[Route('/{todoId}', name: 'update', methods: ['PATCH'], requirements: ['todoId' => '\d+'])]
    #[OA\Parameter(
        name: 'todoId',
        description: 'ID of the todo',
        in: 'path',
        schema: new OA\Schema(type: 'integer', readOnly: true)
    )]
    #[OA\RequestBody(
        description: 'JSON request body.',
        required: true,
        content: new OA\JsonContent(
            ref: new Model(type: ToDo::class, groups: ['update'])
        )
    )]
    #[OA\Response(
        response: 200,
        description: 'Todo updated successfully',
        content: new OA\JsonContent(ref: new Model(type: ToDo::class, groups: ['read']))
    )]
    public function update(
        #[CurrentUser] User $user,
        Request $request,
        int $todoId
    ): Response {
        try {
            $todo = $this->toDoRepository->findByUserAndId($user, $todoId);

            if ($todo === null) {
                throw new HttpException(Response::HTTP_NOT_FOUND, 'Not found');
            }

            $this->serializer->deserialize(
                $request->getContent(),
                ToDo::class,
                'json',
                [
                    'allow_extra_attributes' => false,
                    'groups' => ['update'],
                    AbstractNormalizer::OBJECT_TO_POPULATE => $todo
                ]
            );

            $errors = $this->validator->validate($todo);
            if (count($errors)) {
                $this->throwValidatorException($errors);
            }

            $this->entityManager->flush();
        } catch (DriverException $e) {
            return $this->returnJsonError($e->getMessage(), Response::HTTP_INTERNAL_SERVER_ERROR);
        } catch (Exception $e) {
            return $this->returnJsonError(
                $e->getMessage(),
                $e->getCode() !== 0 ? $e->getCode() : Response::HTTP_UNPROCESSABLE_ENTITY
            );
        }


        return $this->json(
            $todo,
            Response::HTTP_OK,
            [],
            ['groups' => ['read']]
        );
    }

    #[Route('/{todoId}', name: 'delete', methods: ['DELETE'], requirements: ['todoId' => '\d+'])]
    #[OA\Parameter(
        name: 'todoId',
        description: 'ID of the todo',
        in: 'path',
        schema: new OA\Schema(type: 'integer', readOnly: true)
    )]
    public function delete(
        #[CurrentUser] User $user,
        Request $request,
        int $todoId
    ): Response {
        try {
            $todo = $this->toDoRepository->findByUserAndId($user, $todoId);

            if ($todo === null) {
                throw new HttpException(Response::HTTP_NOT_FOUND, 'Not found');
            }

            $this->toDoRepository->remove($todo, true);
        } catch (DriverException $e) {
            return $this->returnJsonError($e->getMessage(), Response::HTTP_INTERNAL_SERVER_ERROR);
        } catch (Exception $e) {
            return $this->returnJsonError(
                $e->getMessage(),
                $e->getCode() !== 0 ? $e->getCode() : Response::HTTP_UNPROCESSABLE_ENTITY
            );
        }


        return $this->json(
            [],
            Response::HTTP_OK,
        );
    }
}
