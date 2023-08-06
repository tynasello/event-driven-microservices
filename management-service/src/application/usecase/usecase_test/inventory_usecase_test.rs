use std::collections::HashMap;

use crate::application::{
    interfaces::i_rest_service::{MockIRestService, RestServiceResponse},
    usecase::{
        add_inventory_usecase::AddInventoryUsecase,
        update_inventory_usecase::UpdateInventoryUsecase,
    },
};

#[tokio::test]
async fn test_adding_inventory() {
    std::env::set_var("INVENTORY_SERVICE_URI", "");
    let mut rest_service_mock = MockIRestService::new();
    rest_service_mock
        .expect_fetch()
        .times(1)
        .returning(|_, _, _, _| {
            Box::pin(async {
                Ok::<RestServiceResponse, String>(RestServiceResponse {
                    body: "body".to_string(),
                    cookies: HashMap::new(),
                })
            })
        });
    let add_inventory_usecase = AddInventoryUsecase::new(&rest_service_mock);
    let result = add_inventory_usecase.execute("", 1).await;
    assert_eq!(result.is_ok(), true);
    assert_eq!(result.unwrap(), "body");
}

#[tokio::test]
async fn test_error_adding_inventory() {
    std::env::set_var("INVENTORY_SERVICE_URI", "");
    let mut rest_service_mock = MockIRestService::new();
    rest_service_mock
        .expect_fetch()
        .times(1)
        .returning(|_, _, _, _| Box::pin(async { Err("error".to_string()) }));
    let add_inventory_usecase = AddInventoryUsecase::new(&rest_service_mock);
    let result = add_inventory_usecase.execute("", 1).await;
    assert_eq!(result.is_err(), true);
    assert_eq!(result.unwrap_err(), "error");
}

#[tokio::test]
async fn test_updating_inventory() {
    std::env::set_var("INVENTORY_SERVICE_URI", "");
    let mut rest_service_mock = MockIRestService::new();
    rest_service_mock
        .expect_fetch()
        .times(1)
        .returning(|_, _, _, _| {
            Box::pin(async {
                Ok::<RestServiceResponse, String>(RestServiceResponse {
                    body: "body".to_string(),
                    cookies: HashMap::new(),
                })
            })
        });
    let update_inventory_usecase = UpdateInventoryUsecase::new(&rest_service_mock);
    let result = update_inventory_usecase.execute("", 1).await;
    assert_eq!(result.is_ok(), true);
    assert_eq!(result.unwrap(), "body");
}

#[tokio::test]
async fn test_error_updating_inventory() {
    std::env::set_var("INVENTORY_SERVICE_URI", "");
    let mut rest_service_mock = MockIRestService::new();
    rest_service_mock
        .expect_fetch()
        .times(1)
        .returning(|_, _, _, _| Box::pin(async { Err("error".to_string()) }));
    let update_inventory_usecase = UpdateInventoryUsecase::new(&rest_service_mock);
    let result = update_inventory_usecase.execute("", 1).await;
    assert_eq!(result.is_err(), true);
    assert_eq!(result.unwrap_err(), "error");
}
