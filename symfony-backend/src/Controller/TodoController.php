<?php

namespace App\Controller;

use App\Repository\ToDoRepository;
use Symfony\Bundle\FrameworkBundle\Controller\AbstractController;
use Symfony\Component\HttpFoundation\Response;
use Symfony\Component\Routing\Attribute\Route;

 #[Route('/api/1.0/todos', name: 'app_todo_')]
final class TodoController extends AbstractController
{
    public function __construct(
        private ToDoRepository $toDoRepository
    ) {}

    #[Route('', name: 'read', methods: ['GET'])]
    public function read(): Response
    {
        $todos = $this->toDoRepository->findByUser($this->getUser());

        return $this->json(
            $todos,
            Response::HTTP_OK,
            [],
            ['groups' => ['read']]
        );
    }
}
