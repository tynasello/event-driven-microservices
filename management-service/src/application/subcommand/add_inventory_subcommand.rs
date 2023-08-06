use clap::ArgMatches;

use crate::application::{
    interfaces::i_rest_service::IRestService, usecase::add_inventory_usecase::AddInventoryUsecase,
};

pub async fn add_inventory_subcommand<'a>(
    add_inventory_matches: &'a ArgMatches<'_>,
    rest_service: &'a dyn IRestService,
) {
    let inventory_label = add_inventory_matches.value_of("label").unwrap();
    let inventory_quantity = add_inventory_matches.value_of("quantity").unwrap();

    let add_inventory_usecase = AddInventoryUsecase::new(rest_service);
    let add_inventory_result = add_inventory_usecase
        .execute(inventory_label, inventory_quantity.parse::<i32>().unwrap())
        .await;
    if let Ok(inventory_id) = add_inventory_result {
        println!("Inventory successfully added: {}", inventory_id);
    } else {
        println!(
            "Error adding inventory: {}",
            add_inventory_result.unwrap_err()
        );
    }
}
