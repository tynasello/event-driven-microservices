use clap::{App, Arg};

pub fn create_subcommand<'a>(name: &'a str, args: &'a Vec<(&'a str, &'a str)>) -> App<'a, 'a> {
    let mut app = App::new(name);

    for (arg_name, arg_help) in args {
        app = app.arg(
            Arg::with_name(arg_name)
                .help(arg_help)
                .required(true)
                .takes_value(true),
        );
    }

    app
}
