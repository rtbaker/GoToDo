<?php


namespace App\Tests\Functional;

use App\Entity\User;
use App\Repository\ToDoRepository;
use App\Repository\UserRepository;
use Symfony\Component\DomCrawler\Crawler;
use Symfony\Bundle\FrameworkBundle\KernelBrowser;
use Symfony\Bundle\FrameworkBundle\Test\WebTestCase;

/**
 * Class BaseWebTestCase
 * @package App\Tests\Functional
 *
 * Use as the base class for all Functional tests to save typing.
 */
class BaseWebTestCase extends WebTestCase
{
    protected ?KernelBrowser $client = null;
    protected ?UserRepository $userRepository = null;
    protected ?ToDoRepository $todoRepository = null;

    protected function setUp(): void
    {
        $kernel = self::bootKernel();
        self::ensureKernelShutdown();
        $this->client = static::createClient();
        $this->userRepository = static::getContainer()->get(UserRepository::class);
        $this->todoRepository = static::getContainer()->get(ToDoRepository::class);
    }

    protected function tearDown(): void
    {
        parent::tearDown();
    }

    public function jsonRequest(
        string $url,
        array $data,
        string $method = 'GET'
    ): Crawler
    {
        return $this->client->request(
                    $method,
                    $url,
                    [],
                    [],
                    ["CONTENT_TYPE" => 'application/json'],
                    json_encode($data)
                );
    }

    public function loginUser(string $email): User
    {
        $user = $this->userRepository->findOneBy(['email' => $email]);

        if ($user === null) {
            throw new \Exception("No such user in test: " . $email);
        }

        $this->client->loginUser($user);

        return $user;
    }
}