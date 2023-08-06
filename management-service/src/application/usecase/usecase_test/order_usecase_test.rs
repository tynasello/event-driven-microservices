use std::collections::HashMap;

use crate::application::{
    interfaces::i_rest_service::{MockIRestService, RestServiceResponse},
    usecase::{
        create_order_usecase::CreateOrderUsecase, get_order_status_usecase::GetOrderStatusUsecase,
    },
};

#[tokio::test]
async fn test_creating_order() {
    std::env::set_var("ORDER_SERVICE_URI", "");
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
    let create_order_usecase = CreateOrderUsecase::new(&rest_service_mock);
    let result = create_order_usecase.execute("".to_string(), "", 1).await;
    assert_eq!(result.is_ok(), true);
    assert_eq!(result.unwrap(), "body");
}

#[tokio::test]
async fn test_error_creating_order() {
    std::env::set_var("ORDER_SERVICE_URI", "");
    let mut rest_service_mock = MockIRestService::new();
    rest_service_mock
        .expect_fetch()
        .times(1)
        .returning(|_, _, _, _| Box::pin(async { Err("error".to_string()) }));
    let create_order_usecase = CreateOrderUsecase::new(&rest_service_mock);
    let result = create_order_usecase.execute("".to_string(), "", 1).await;
    assert_eq!(result.is_err(), true);
    assert_eq!(result.unwrap_err(), "error");
}

#[tokio::test]
async fn test_get_order_status() {
    std::env::set_var("ORDER_SERVICE_URI", "");
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
    let get_order_status_usecase = GetOrderStatusUsecase::new(&rest_service_mock);
    let result = get_order_status_usecase.execute("".to_string(), 1).await;
    assert_eq!(result.is_ok(), true);
    assert_eq!(result.unwrap(), "body");
}

#[tokio::test]
async fn test_error_getting_order_status() {
    std::env::set_var("ORDER_SERVICE_URI", "");
    let mut rest_service_mock = MockIRestService::new();
    rest_service_mock
        .expect_fetch()
        .times(1)
        .returning(|_, _, _, _| Box::pin(async { Err("error".to_string()) }));
    let get_order_status_usecase = GetOrderStatusUsecase::new(&rest_service_mock);
    let result = get_order_status_usecase.execute("".to_string(), 1).await;
    assert_eq!(result.is_err(), true);
    assert_eq!(result.unwrap_err(), "error");
}
