package com.example.orderservice.infra.services.jparepository;

import com.example.orderservice.infra.persistedmodel.OrderModel;
import org.springframework.data.jpa.repository.JpaRepository;

public interface JpaOrderRepository extends JpaRepository<OrderModel, Long> {}
