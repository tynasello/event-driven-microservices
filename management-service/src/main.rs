use application::{
    interfaces::i_rest_service::IRestService,
    subcommand::{
        add_inventory_subcommand::add_inventory_subcommand,
        create_order_subcommand::create_order_subcommand,
        get_order_status_subcommand::get_order_status_subcommand,
        login_subcommand::login_subcommand, signup_subcommand::signup_subcommand,
        subcommand_helper::create_subcommand,
        update_inventory_subcommand::update_inventory_subcommand,
    },
};
use clap::App;
use domain::cli::cli_commands::get_cli_commands;

use crate::{
    application::interfaces::i_message_broker_consumer_service::IMessageBrokerConsumerService,
    infra::{
        message_broker::kafka_config::KafkaConsumer,
        service::{
            message_broker_consumer_service::MessageBrokerConsumerService,
            rest_service::RestService,
        },
    },
};

mod application;
mod domain;
mod infra;

#[tokio::main]
async fn main() {
    dotenv::dotenv().ok();

    let cli_subcommands = get_cli_commands();

    let matches = App::new("Edrims Management Service")
        .version("1.0")
        .author("tynasello")
        .about("A CLI application to interact with and oversee the edrims system")
        .subcommands(
            cli_subcommands
                .iter()
                .map(|&(name, ref args)| create_subcommand(name, args))
                .collect::<Vec<_>>(),
        )
        .get_matches();

    let rest_service: Box<dyn IRestService> = Box::new(RestService {});

    match matches.subcommand() {
        ("login", Some(login_matches)) => {
            login_subcommand(login_matches, &rest_service).await;
        }
        ("signup", Some(signup_matches)) => {
            signup_subcommand(signup_matches, &rest_service).await;
        }
        ("create-order", Some(create_order_matches)) => {
            create_order_subcommand(create_order_matches, &rest_service).await;
        }
        ("get-order-status", Some(get_order_status_matches)) => {
            get_order_status_subcommand(get_order_status_matches, &rest_service).await;
        }
        ("add-inventory", Some(add_inventory_matches)) => {
            add_inventory_subcommand(add_inventory_matches, &rest_service).await;
        }
        ("update-inventory", Some(update_inventory_matches)) => {
            update_inventory_subcommand(update_inventory_matches, &rest_service).await;
        }
        _ => {
            println!("Invalid command. Use --help to see the available commands.");
        }
    }

    // // consume kafka events
    // let kafka_consumer = &mut KafkaConsumer::new();
    // let message_broker_consumer_service = &mut MessageBrokerConsumerService { kafka_consumer };
    // message_broker_consumer_service.start_consuming();

    return;
}
