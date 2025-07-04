<?php

namespace App\DataFixtures;

use App\Entity\User;
use Doctrine\Bundle\FixturesBundle\Fixture;
use Doctrine\Persistence\ObjectManager;
use Symfony\Component\PasswordHasher\Hasher\UserPasswordHasherInterface;

class UserFixtures extends Fixture
{
    public function __construct(private UserPasswordHasherInterface $passwordHasher) {}

    public function load(ObjectManager $manager): void
    {
        $user1 = new User();
        $user1->setEmail('test1@test.com');
        $user1->setPassword($this->passwordHasher->hashPassword($user1, 'password'));
        $manager->persist($user1);

        $user2 = new User();
        $user2->setEmail('test2@test.com');
        $user2->setPassword($this->passwordHasher->hashPassword($user2, 'password'));
        $manager->persist($user2);

        $user3 = new User();
        $user3->setEmail('test3@test.com');
        $user3->setPassword($this->passwordHasher->hashPassword($user3, 'password'));
        $manager->persist($user3);

        $manager->flush();

        $this->addReference('user1', $user1);
        $this->addReference('user2', $user2);
        $this->addReference('user3', $user3);    
    }
}
