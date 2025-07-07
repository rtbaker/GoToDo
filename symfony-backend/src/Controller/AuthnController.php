<?php

declare(strict_types=1);

namespace App\Controller;

use Symfony\Bundle\FrameworkBundle\Controller\AbstractController;
use Symfony\Component\HttpFoundation\Response;
use Symfony\Component\Routing\Attribute\Route;
use OpenApi\Attributes as OA;
use Nelmio\ApiDocBundle\Attribute\Model;

final class AuthnController extends AbstractController
{
    #[Route('/api/1.0/login', name: 'api_login', methods: ['POST'])]
    #[OA\RequestBody(
        description: 'JSON request body.',
        required: true,
        content: new OA\JsonContent(
            type: 'object',
            properties: [
                    new OA\Property(property: "email", type: "string", example: "test@test.com"),
                    new OA\Property(property: "password", type: "string", example: "my-password"),
            ],
            required: ['email', 'password']
        )
    )]
    public function index(): Response
    {
        return $this->json(
            $this->getUser(),
            Response::HTTP_OK,
            [],
            ['groups' => ['read']]
        );
    }
}
