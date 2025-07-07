<?php

declare(strict_types=1);

namespace App\Controller;

use Exception;
use App\Entity\ToDo;
use App\Entity\User;
use App\Repository\ToDoRepository;
use Doctrine\DBAL\Exception\DriverException;
use Doctrine\ORM\EntityManagerInterface;
use Symfony\Component\HttpFoundation\Request;
use Symfony\Component\HttpFoundation\Response;
use Symfony\Component\Routing\Attribute\Route;
use Symfony\Component\Serializer\SerializerInterface;
use Symfony\Component\Security\Http\Attribute\CurrentUser;
use Symfony\Bundle\FrameworkBundle\Controller\AbstractController;
use Symfony\Component\HttpKernel\Exception\HttpException;
use Symfony\Component\Serializer\Normalizer\AbstractNormalizer;

 #[Route('/api/1.0/todos', name: 'app_todo_')]
final class TodoController extends AbstractController
{
    use ErrorTrait;

    public function __construct(
        private ToDoRepository $toDoRepository,
        private SerializerInterface $serializer,
        private EntityManagerInterface $entityManager
    ) {
    }

    #[Route('', name: 'read', methods: ['GET'])]
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
