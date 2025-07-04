<?php

namespace App\Controller;

use Symfony\Bundle\FrameworkBundle\Controller\AbstractController;
use Symfony\Component\HttpFoundation\JsonResponse;
use Symfony\Component\Routing\Attribute\Route;

final class UptestController extends AbstractController
{
    #[Route('/uptest', name: 'app_uptest')]
    public function index(): JsonResponse
    {
        return $this->json([

            'status' => 'ok'
        ]);
    }
}
