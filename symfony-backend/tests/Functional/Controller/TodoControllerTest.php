<?php

namespace App\Tests\Functional\Controller;

use App\Tests\Functional\BaseWebTestCase;
use Symfony\Component\HttpFoundation\Response;
use Symfony\Bundle\FrameworkBundle\Test\WebTestCase;

class TodoControllerTest extends BaseWebTestCase
{
    const TODO_API = '/api/1.0/todos';

    public function testGetTodos(): void
    {
        $this->loginUser('test1@test.com');

        $req = $this->jsonRequest(
            self::TODO_API,
            ['email' => 'test1@test.com', 'password' => 'password'],
            'GET'
        );

        $this->assertEquals(Response::HTTP_OK, $this->client->getResponse()->getStatusCode());

        $body = json_decode($this->client->getResponse()->getContent());
        
        // User test1@test.com has 4 todos
        $this->assertEquals(4, count($body));
    }

    public function testCanDeleteTodo(): void
    {
        $user = $this->loginUser('test1@test.com');
        $todos = $this->todoRepository->findByUser($user);

        $this->assertGreaterThan(0, count($todos));

        $todoToDelete = $todos[0];

        $req = $this->jsonRequest(
            sprintf("%s/%d", self::TODO_API, $todoToDelete->getId()),
            [],
            'DELETE'
        );

        $this->assertEquals(Response::HTTP_OK, $this->client->getResponse()->getStatusCode());
    }

    public function testCantDeleteSomeoneElsesTodo(): void
    {
        $this->loginUser('test1@test.com');

        $user = $this->userRepository->findOneBy(['email' => 'test2@test.com']);
        $todos = $this->todoRepository->findByUser($user);

        $this->assertGreaterThan(0, count($todos));

        $todoToDelete = $todos[0];

        $req = $this->jsonRequest(
            sprintf("%s/%d", self::TODO_API, $todoToDelete->getId()),
            [],
            'DELETE'
        );

        $this->assertNotEquals(Response::HTTP_OK, $this->client->getResponse()->getStatusCode());
    }

    public function testCantUpdateSomeoneElsesTodo(): void
        {
            $this->loginUser('test1@test.com');
    
            $user = $this->userRepository->findOneBy(['email' => 'test2@test.com']);
            $todos = $this->todoRepository->findByUser($user);
    
            $this->assertGreaterThan(0, count($todos));
    
            $todoToDelete = $todos[0];
    
            $req = $this->jsonRequest(
                sprintf("%s/%d", self::TODO_API, $todoToDelete->getId()),
                ['title' => 'new title'],
                'PATCH'
            );
    
            $this->assertNotEquals(Response::HTTP_OK, $this->client->getResponse()->getStatusCode());
        }
}
