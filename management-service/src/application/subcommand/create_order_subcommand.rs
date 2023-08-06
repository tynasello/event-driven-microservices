use clap::ArgMatches;

use crate::application::{
    interfaces::i_rest_service::IRestService, usecase::create_order_usecase::CreateOrderUsecase,
};

pub async fn create_order_subcommand<'a>(
    create_order_matches: &'a ArgMatches<'_>,
    rest_service: &'a dyn IRestService,
) {
    let access_token = create_order_matches.value_of("access-token").unwrap();
    let product_name = create_order_matches.value_of("product-name").unwrap();
    let product_quantity = create_order_matches.value_of("product-quantity").unwrap();

    let create_order_usecase = CreateOrderUsecase::new(rest_service);
    let create_order_result = create_order_usecase
        .execute(access_token, product_name, product_quantity)
        .await;

    if let Ok(order) = create_order_result {
        println!("Order successfully created: {}", order);
    } else {
        println!("Error creating order: {}", create_order_result.unwrap_err());
    }
}
