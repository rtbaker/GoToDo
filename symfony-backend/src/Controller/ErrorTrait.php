<?php

declare(strict_types=1);

namespace App\Controller;

use Symfony\Component\HttpFoundation\JsonResponse;

trait ErrorTrait
{
    public function returnJsonError(
        string $message,
        int $statusCode = 500
    ): JsonResponse {
        return $this->json(
            [
                'code' => $statusCode,
                'message' => $message,
            ],
            $statusCode
        );
    }
}
