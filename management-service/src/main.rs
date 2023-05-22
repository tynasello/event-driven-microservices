use application::{
    interfaces::i_rest_service::IRestService,
    usecase::{login_usecase::LoginUsecase, update_inventory_usecase::UpdateInventoryUsecase},
};

use crate::{
    application::{
        interfaces::i_message_broker_consumer_service::IMessageBrokerConsumerService,
        usecase::{
            add_inventory_usecase::AddInventoryUsecase, create_order_usecase::CreateOrderUsecase,
            get_order_status_usecase::GetOrderStatusUsecase, signup_usecase::SignupUsecase,
        },
    },
    infra::{
        message_broker::kafka_config::KafkaConsumer,
        service::{
            message_broker_consumer_service::MessageBrokerConsumerService,
            rest_service::RestService,
        },
    },
};

mod application;
mod infra;

#[tokio::main]
async fn main() {
    dotenv::dotenv().ok();

    // consume kafka events
    let kafka_consumer = &mut KafkaConsumer::new();
    let message_broker_consumer_service = &mut MessageBrokerConsumerService { kafka_consumer };
    message_broker_consumer_service.start_consuming();

    let rest_service: Box<dyn IRestService> = Box::new(RestService {});

    // login / signup
    let signup_user_service = SignupUsecase::new(&rest_service);
    let _signup_result = signup_user_service.execute("baba", "jaga");

    let login_user_service = LoginUsecase::new(&rest_service);
    let access_token = login_user_service.execute("baba", "jaga").await;

    // create order
    match access_token {
        Ok(access_token) => {
            println!("Access token: {}", access_token);
            let create_order_usecase = CreateOrderUsecase::new(&rest_service);
            let create_order_result = create_order_usecase
                .execute(access_token, "carrot", 10)
                .await;
            return println!(
                "Result: {}",
                create_order_result.unwrap_or_else(|error| error)
            );
        }
        Err(_) => {
            println!("Error creating a order");
        }
    }

    // get order service
    match access_token {
        Ok(access_token) => {
            println!("Access token: {}", access_token);
            let get_order_status_usecase = GetOrderStatusUsecase::new(&rest_service);
            let get_order_status_result = get_order_status_usecase.execute(access_token, 1).await;
            return println!(
                "Result: {}",
                get_order_status_result.unwrap_or_else(|error| error)
            );
        }
        Err(_) => {
            println!("Error getting order");
        }
    }

    // add inventory
    let add_inventory_usecase = AddInventoryUsecase::new(&rest_service);
    let add_inventory_result = add_inventory_usecase.execute("carrot", 10).await;
    println!(
        "Result: {}",
        add_inventory_result.unwrap_or_else(|error| error)
    );

    // update inventory
    let update_inventory_usecase = UpdateInventoryUsecase::new(&rest_service);
    let update_inventory_result = update_inventory_usecase.execute("carrot", 10).await;
    println!(
        "Result: {}",
        update_inventory_result.unwrap_or_else(|error| error)
    );

    return;
}
