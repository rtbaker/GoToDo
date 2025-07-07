<?php

declare(strict_types=1);

namespace App\Controller;

use Symfony\Component\HttpFoundation\Response;
use Symfony\Component\HttpFoundation\JsonResponse;
use Symfony\Component\HttpKernel\Exception\HttpException;
use Symfony\Component\Validator\ConstraintViolationListInterface;

trait ErrorTrait
{
    /**
     * Format a standard (for this project) error message.
     */
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

    /**
     * Turn a list of validator exceptions into an Exception
     */
    protected function throwValidatorException(
        ConstraintViolationListInterface $errors,
        int $code = Response::HTTP_UNPROCESSABLE_ENTITY
    ): void {
        $errorsString = implode(
            "\n",
            array_map(
                fn($item) => $item->getPropertyPath() . ': ' . $item->getMessage(),
                iterator_to_array($errors)
            )
        );
        throw new HttpException($code, $errorsString, null, [], $code);
    }
}
