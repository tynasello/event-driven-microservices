use clap::ArgMatches;

use crate::application::{
    interfaces::i_rest_service::IRestService,
    usecase::get_order_status_usecase::GetOrderStatusUsecase,
};

pub async fn get_order_status_subcommand<'a>(
    get_order_status_matches: &'a ArgMatches<'_>,
    rest_service: &'a dyn IRestService,
) {
    let access_token = get_order_status_matches.value_of("access-token").unwrap();
    let order_id = get_order_status_matches.value_of("order-id").unwrap();

    let get_order_status_usecase = GetOrderStatusUsecase::new(rest_service);
    let get_order_status_result = get_order_status_usecase
        .execute(access_token.to_string(), order_id.parse::<i32>().unwrap())
        .await;

    if let Ok(order_status) = get_order_status_result {
        println!("Order status: {}", order_status);
    } else {
        println!(
            "Error getting order status: {}",
            get_order_status_result.unwrap_err()
        );
    }
}
