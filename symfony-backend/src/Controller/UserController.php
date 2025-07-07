<?php

declare(strict_types=1);

namespace App\Controller;

use Symfony\Bundle\FrameworkBundle\Controller\AbstractController;
use Symfony\Component\HttpFoundation\Response;
use Symfony\Component\Routing\Attribute\Route;

final class UserController extends AbstractController
{
    #[Route('/api/1.0/user', name: 'app_user')]
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
