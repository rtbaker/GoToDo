<?php

namespace App\Tests\Functional\Controller;

use App\Tests\Functional\BaseWebTestCase;
use Symfony\Component\HttpFoundation\Response;
use Symfony\Bundle\FrameworkBundle\Test\WebTestCase;

class AuthnControllerTest extends BaseWebTestCase
{
    public function testGoodUser(): void
    {
        $req = $this->jsonRequest(
            '/api/1.0/login',
            ['email' => 'test1@test.com', 'password' => 'password'],
            'POST'
        );
        $this->assertEquals(Response::HTTP_OK, $this->client->getResponse()->getStatusCode());
    }

    public function testBadUser(): void
    {
        $req = $this->jsonRequest(
                    '/api/1.0/login',
                    ['email' => 'test1@test.com', 'password' => 'NOT THIS'],
                    'POST'
                );

        $this->assertEquals(Response::HTTP_UNAUTHORIZED, $this->client->getResponse()->getStatusCode());
    }
}
