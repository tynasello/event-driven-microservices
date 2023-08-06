use std::collections::HashMap;

use crate::application::interfaces::i_rest_service::IRestService;
use crate::infra::service::rest_service::RestService;

#[tokio::test]
async fn test_get() {
    let mut user_service_mock = mockito::Server::new();

    let mock_endpoint = "/hello";

    user_service_mock
        .mock("GET", mock_endpoint)
        .with_status(201)
        .create();
    let user_service_mock_url = user_service_mock.url();

    let reqwest_rest_service = RestService::new(None);

    let response = reqwest_rest_service
        .fetch(
            "GET",
            &format!("{}{}", user_service_mock_url, mock_endpoint),
            "",
            HashMap::new(),
        )
        .await;

    assert!(response.is_ok());
}

#[tokio::test]
async fn test_post() {
    let mut user_service_mock = mockito::Server::new();

    let mock_endpoint = "/hello";

    user_service_mock
        .mock("POST", mock_endpoint)
        .with_status(201)
        .create();
    let user_service_mock_url = user_service_mock.url();

    let reqwest_rest_service = RestService::new(None);

    let response = reqwest_rest_service
        .fetch(
            "POST",
            &format!("{}{}", user_service_mock_url, mock_endpoint),
            "",
            HashMap::new(),
        )
        .await;

    assert!(response.is_ok());
}

#[tokio::test]
async fn test_get_with_invalid_method() {
    let reqwest_rest_service = RestService::new(None);

    let response = reqwest_rest_service
        .fetch("INVALID_METHOD", "", "", HashMap::new())
        .await;

    assert!(response.is_err());
}

#[tokio::test]
async fn test_get_with_invalid_endpoint() {
    let reqwest_rest_service = RestService::new(None);

    let response = reqwest_rest_service
        .fetch("GET", "", "", HashMap::new())
        .await;

    assert!(response.is_err());
}

#[tokio::test]
async fn test_get_failed_response() {
    let mut user_service_mock = mockito::Server::new();

    let mock_endpoint = "/hello";

    user_service_mock
        .mock("GET", mock_endpoint)
        .with_status(400)
        .create();
    let user_service_mock_url = user_service_mock.url();

    let reqwest_rest_service = RestService::new(None);

    let response = reqwest_rest_service
        .fetch(
            "GET",
            &format!("{}{}", user_service_mock_url, mock_endpoint),
            "",
            HashMap::new(),
        )
        .await;

    assert!(response.is_err());
}

#[tokio::test]
async fn test_gets_response_body() {
    let mut user_service_mock = mockito::Server::new();

    let mock_endpoint = "/hello";
    let mock_response_body = "world";

    user_service_mock
        .mock("GET", mock_endpoint)
        .with_status(201)
        .with_body(mock_response_body)
        .create();
    let user_service_mock_url = user_service_mock.url();

    let reqwest_rest_service = RestService::new(None);

    let response = reqwest_rest_service
        .fetch(
            "GET",
            &format!("{}{}", user_service_mock_url, mock_endpoint),
            "",
            HashMap::new(),
        )
        .await;

    assert!(response.is_ok());
    assert_eq!(response.unwrap().body, mock_response_body);
}

#[tokio::test]
async fn test_gets_access_token_from_response_cookies() {
    let mut user_service_mock = mockito::Server::new();

    let mock_endpoint = "/hello";
    let mock_response_access_token = "at";

    user_service_mock
        .mock("GET", mock_endpoint)
        .with_status(201)
        .with_header(
            "Set-Cookie",
            &format!("access-token={}", mock_response_access_token),
        )
        .create();
    let user_service_mock_url = user_service_mock.url();

    let reqwest_rest_service = RestService::new(None);

    let response = reqwest_rest_service
        .fetch(
            "GET",
            &format!("{}{}", user_service_mock_url, mock_endpoint),
            "",
            HashMap::new(),
        )
        .await;

    let expected_response_cookies: HashMap<String, String> = HashMap::from_iter(vec![(
        "access-token".to_string(),
        mock_response_access_token.to_string(),
    )]);

    assert!(response.is_ok());
    for (key, value) in response.unwrap().cookies.iter() {
        assert_eq!(expected_response_cookies.get(key), Some(value));
    }
}
