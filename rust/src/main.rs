#[macro_use]
extern crate diesel;
#[macro_use]
extern crate log;

mod cmd;
mod db;
mod models;
mod schema;

use cmd::*;
use diesel::result::Error;
use dotenv::dotenv;
use structopt::StructOpt;

fn main() -> Result<(), Error> {
    dotenv().ok();
    env_logger::init();
    let conn = db::establish_connection();
    let opt = Opt::from_args();

    match opt.cmd {
        Command::Select { table } => match table {
            TargetTable::Person => select_person(&conn),
            TargetTable::Post => select_post(&conn),
            TargetTable::PersonWithPost { num } => select_person_with_post(&conn, num),
        },
        Command::Join => join(&conn),
        Command::Insert {
            num_person,
            num_posts_per_person,
        } => insert(&conn, num_person, num_posts_per_person),
        Command::Delete => delete(&conn),
    }
}
