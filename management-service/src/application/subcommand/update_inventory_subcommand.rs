use clap::ArgMatches;

use crate::application::{
    interfaces::i_rest_service::IRestService,
    usecase::update_inventory_usecase::UpdateInventoryUsecase,
};

pub async fn update_inventory_subcommand<'a>(
    update_inventory_matches: &'a ArgMatches<'_>,
    rest_service: &'a dyn IRestService,
) {
    let inventory_label = update_inventory_matches.value_of("label").unwrap();
    let quantity_to_add = update_inventory_matches
        .value_of("quantity-to-add")
        .unwrap();

    let update_inventory_usecase = UpdateInventoryUsecase::new(rest_service);
    let update_inventory_result = update_inventory_usecase
        .execute(inventory_label, quantity_to_add.parse::<i32>().unwrap())
        .await;
    if let Ok(inventory_id) = update_inventory_result {
        println!("Inventory successfully updated: {}", inventory_id);
    } else {
        println!(
            "Error updating inventory: {}",
            update_inventory_result.unwrap_err()
        );
    }
}
