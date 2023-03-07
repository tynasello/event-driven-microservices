package com.example.orderservice.infra.repository;

import com.example.orderservice.infra.persistedmodel.OrderModel;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

@Repository
public interface OrderRepository extends JpaRepository<OrderModel, Long> {}
