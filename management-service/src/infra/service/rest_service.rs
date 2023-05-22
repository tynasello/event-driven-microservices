use std::collections::HashMap;

use async_trait::async_trait;
use reqwest::header::{HeaderMap, HeaderValue, COOKIE};

use crate::application::interfaces::i_rest_service::{IRestService, RestServiceResponse};

pub struct RestService {}

#[async_trait]
impl IRestService for RestService {
    async fn fetch(
        &self,
        endpoint: &str,
        access_token: &str,
        body: HashMap<&str, &str>,
    ) -> Result<RestServiceResponse, String> {
        let client = reqwest::Client::new();

        let mut request_headers = HeaderMap::new();
        let access_token_cookie: &str = &["access_token=", access_token].concat();
        request_headers.insert(COOKIE, HeaderValue::from_str(access_token_cookie).unwrap());

        let response = client
            .post(endpoint)
            .headers(request_headers)
            .json(&body)
            .send()
            .await;

        match response {
            Ok(response) => {
                if response.status().is_success() == false {
                    // instead return actual api response
                    return Err("Error making api call".to_string());
                }

                let response_headers = response.headers();
                let raw_cookies = response_headers
                    .get("set-cookie")
                    .unwrap()
                    .to_str()
                    .unwrap();

                let access_token = raw_cookies
                    .split(';')
                    .find(|s| s.starts_with("access-token="))
                    .map(|s| s.split('=').skip(1).next().unwrap_or(""))
                    .map(|s| s.trim())
                    .unwrap_or("");

                let mut response_cookies = HashMap::new();
                response_cookies.insert("access_token".to_string(), access_token.to_string());

                let response_body = response.text().await.unwrap_or_else(|_| "".to_string());

                let rest_service_response = RestServiceResponse {
                    body: response_body,
                    cookies: response_cookies,
                };

                return Ok(rest_service_response);
            }
            Err(error) => {
                return Err(error.to_string());
            }
        }
    }
}
