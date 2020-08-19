use super::models::*;
use super::schema::person::dsl::*;
use super::schema::post::dsl::*;
use diesel::prelude::*;
use diesel::result::Error;
use structopt::StructOpt;

#[derive(Debug, StructOpt)]
pub struct Opt {
    #[structopt(subcommand)]
    pub cmd: Command,
}

#[derive(Debug, StructOpt)]
pub enum Command {
    Select {
        #[structopt(subcommand)]
        table: TargetTable,
    },
    Join,
    Delete,
    Insert {
        num_person: i64,
        num_posts_per_person: i64,
    },
}

#[derive(Debug, StructOpt)]
pub enum TargetTable {
    Person,
    Post,
    PersonWithPost { num: i64 },
}

pub fn select_person(conn: &PgConnection) -> Result<(), Error> {
    let results = person.load::<Person>(conn).unwrap();

    for r in results {
        info!("{:?}", r);
    }

    Ok(())
}

pub fn select_post(conn: &PgConnection) -> Result<(), Error> {
    let results = post.load::<Post>(conn).unwrap();

    for r in results {
        info!("{:?}", r);
    }

    Ok(())
}

pub fn select_person_with_post(conn: &PgConnection, num: i64) -> Result<(), Error> {
    let persons = person.limit(num).load::<Person>(conn).unwrap();
    let posts: Vec<Vec<Post>> = Post::belonging_to(&persons)
        .load::<Post>(conn)?
        .grouped_by(&persons);

    let results = persons.into_iter().zip(posts).collect::<Vec<_>>();

    for r in results {
        info!("{:?}", r);
    }

    Ok(())
}

pub fn join(conn: &PgConnection) -> Result<(), Error> {
    let results = post
        .inner_join(person)
        .load::<PostWithPersonTuple>(conn)
        .unwrap();

    for r in results {
        info!("{:?}", r.flatten());
    }

    Ok(())
}

pub fn insert(
    conn: &PgConnection,
    num_person: i64,
    num_posts_per_person: i64,
) -> Result<(), Error> {
    let new_persons: Vec<Person> = (0..num_person)
        .map(|x| Person::new(format!("hoge_{0}", x)))
        .collect();

    let new_posts: Vec<Post> = new_persons
        .iter()
        .enumerate()
        .map::<Vec<Post>, _>(|(i, p)| {
            (0..num_posts_per_person)
                .map(|j| Post::new(p.id, format!("fuga_{0}_{1}", i, j)))
                .collect()
        })
        .flatten()
        .collect();

    conn.transaction::<(), Error, _>(|| {
        diesel::insert_into(person)
            .values(new_persons)
            .execute(conn)
            .unwrap();

        diesel::insert_into(post)
            .values(new_posts)
            .execute(conn)
            .unwrap();

        Ok(())
    })
}

pub fn delete(conn: &PgConnection) -> Result<(), Error> {
    conn.transaction::<(), Error, _>(|| {
        diesel::delete(post).execute(conn).unwrap();
        diesel::delete(person).execute(conn).unwrap();

        Ok(())
    })
}
