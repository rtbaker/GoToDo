<?php

namespace App\DataFixtures;

use App\Entity\ToDo;
use App\Entity\User;
use Doctrine\Bundle\FixturesBundle\Fixture;
use Doctrine\Persistence\ObjectManager;
use Doctrine\Common\DataFixtures\DependentFixtureInterface;

class TodoFixtures extends Fixture implements DependentFixtureInterface
{
    public function load(ObjectManager $manager): void
    {
        $user1 = $this->getReference("user1", User::class);
        $user2 = $this->getReference("user2", User::class);
        $user3 = $this->getReference("user3", User::class);

        $todos = [
            [$user1, 'todo1', 'todo1 description', 1, false],
            [$user1, 'todo2', 'todo2 description', 2, false],
            [$user1, 'todo3', 'todo3 description', 3, false],
            [$user1, 'todo4', 'todo4 description', 1, false],
            [$user2, 'todo5', 'todo5 description', 1, false],
            [$user3, 'todo6', 'todo6 description', 1, false],
        ];

        foreach ($todos as $todo) {
            $newTodo = new ToDo();
            $newTodo->setUser($todo[0]);
            $newTodo->setTitle($todo[1]);
            $newTodo->setDescription($todo[2]);
            $newTodo->setPriority($todo[3]);
            $newTodo->setCompleted($todo[4]);

            $manager->persist($newTodo);
        }

        $manager->flush();
    }

    public function getDependencies(): array
    {
        return [
            UserFixtures::class,
        ];
    }
}
