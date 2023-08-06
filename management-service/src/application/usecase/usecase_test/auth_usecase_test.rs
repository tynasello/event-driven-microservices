use std::collections::HashMap;

use crate::application::{
    interfaces::i_rest_service::{MockIRestService, RestServiceResponse},
    usecase::{login_usecase::LoginUsecase, signup_usecase::SignupUsecase},
};

#[tokio::test]
async fn test_signup() {
    std::env::set_var("USER_SERVICE_URI", "");

    let mut rest_service_mock = MockIRestService::new();

    rest_service_mock
        .expect_fetch()
        .times(1)
        .returning(|_, _, _, _| {
            let mut response_cookies = HashMap::new();
            response_cookies.insert("access-token".to_string(), "value".to_string());
            Box::pin(async {
                Ok::<RestServiceResponse, String>(RestServiceResponse {
                    body: "body".to_string(),
                    cookies: response_cookies,
                })
            })
        });
    let signup_usecase = SignupUsecase::new(&rest_service_mock);
    let result = signup_usecase.execute("", "").await;
    assert_eq!(result.is_ok(), true);
    assert_eq!(result.unwrap(), "value");
}

#[tokio::test]
async fn test_error_signup() {
    std::env::set_var("USER_SERVICE_URI", "");
    let mut rest_service_mock = MockIRestService::new();
    rest_service_mock
        .expect_fetch()
        .times(1)
        .returning(|_, _, _, _| Box::pin(async { Err("error".to_string()) }));
    let signup_usecase = SignupUsecase::new(&rest_service_mock);
    let result = signup_usecase.execute("", "").await;
    assert_eq!(result.is_err(), true);
    assert_eq!(result.unwrap_err(), "error");
}

#[tokio::test]
async fn test_login() {
    std::env::set_var("USER_SERVICE_URI", "");

    let mut rest_service_mock = MockIRestService::new();
    rest_service_mock
        .expect_fetch()
        .times(1)
        .returning(|_, _, _, _| {
            let mut response_cookies = HashMap::new();
            response_cookies.insert("access-token".to_string(), "value".to_string());
            Box::pin(async {
                Ok::<RestServiceResponse, String>(RestServiceResponse {
                    body: "body".to_string(),
                    cookies: response_cookies,
                })
            })
        });
    let login_usecase = LoginUsecase::new(&rest_service_mock);
    let result = login_usecase.execute("", "").await;
    assert_eq!(result.is_ok(), true);
    assert_eq!(result.unwrap(), "value");
}

#[tokio::test]
async fn test_error_login() {
    std::env::set_var("USER_SERVICE_URI", "");
    let mut rest_service_mock = MockIRestService::new();
    rest_service_mock
        .expect_fetch()
        .times(1)
        .returning(|_, _, _, _| Box::pin(async { Err("error".to_string()) }));
    let login_usecase = LoginUsecase::new(&rest_service_mock);
    let result = login_usecase.execute("", "").await;
    assert_eq!(result.is_err(), true);
    assert_eq!(result.unwrap_err(), "error");
}
