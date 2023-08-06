use clap::ArgMatches;

use crate::application::{
    interfaces::i_rest_service::IRestService, usecase::signup_usecase::SignupUsecase,
};

pub async fn signup_subcommand<'a>(
    signup_matches: &'a ArgMatches<'_>,
    rest_service: &'a dyn IRestService,
) {
    let username = signup_matches.value_of("username").unwrap();
    let password = signup_matches.value_of("password").unwrap();

    let signup_user_service = SignupUsecase::new(rest_service);
    let signup_user_result = signup_user_service.execute(username, password).await;
    if let Ok(access_token) = signup_user_result {
        println!("Access token: {}", access_token);
    } else {
        println!("Error signup up: {}", signup_user_result.unwrap_err());
    }
}
