<?php

declare(strict_types=1);

namespace App\Repository;

use App\Entity\ToDo;
use Doctrine\Bundle\DoctrineBundle\Repository\ServiceEntityRepository;
use Doctrine\Persistence\ManagerRegistry;
use Symfony\Component\Security\Core\User\UserInterface;

/**
 * @extends ServiceEntityRepository<ToDo>
 */
class ToDoRepository extends ServiceEntityRepository
{
    public function __construct(ManagerRegistry $registry)
    {
        parent::__construct($registry, ToDo::class);
    }

    public function add(Todo $entity, bool $flush = false): void
    {
        $this->getEntityManager()->persist($entity);

        if ($flush) {
            $this->getEntityManager()->flush();
        }
    }

    public function remove(Todo $entity, bool $flush = false): void
    {
        $this->getEntityManager()->remove($entity);

        if ($flush) {
            $this->getEntityManager()->flush();
        }
    }

    /**
     * @return ToDo[] Returns an array of ToDo objects
     */
    public function findByUser(UserInterface $user): array
    {
        /** @var Todo[] */
        $todos = $this->createQueryBuilder('t')
            ->andWhere('t.user = :user')
            ->setParameter('user', $user)
            ->orderBy('t.id', 'ASC')
            ->getQuery()
            ->getResult()
        ;

        return $todos;
    }

    /**
     * @return ?ToDo Returns an array of ToDo objects
     */
    public function findByUserAndId(UserInterface $user, int $id): ?ToDo
    {
        /** @var ?Todo */
        $todo = $this->createQueryBuilder('t')
            ->andWhere('t.id = :id')
            ->andWhere('t.user = :user')
            ->setParameter('id', $id)
            ->setParameter('user', $user)
            ->getQuery()
            ->getOneOrNullResult()
        ;

        return $todo;
    }

    //    public function findOneBySomeField($value): ?ToDo
    //    {
    //        return $this->createQueryBuilder('t')
    //            ->andWhere('t.exampleField = :val')
    //            ->setParameter('val', $value)
    //            ->getQuery()
    //            ->getOneOrNullResult()
    //        ;
    //    }
}
