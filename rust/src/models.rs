use crate::schema::person;
use crate::schema::post;
use chrono::{NaiveDateTime, Utc};
use uuid::Uuid;

#[derive(Debug, PartialEq, Identifiable, Queryable, Insertable)]
#[table_name = "person"]
pub struct Person {
    pub id: Uuid,
    pub name: String,
    pub create_at: NaiveDateTime,
    pub update_at: NaiveDateTime,
}

#[derive(Debug, Queryable, Insertable)]
#[table_name = "person"]
pub struct NewPerson {
    pub name: String,
}

#[derive(Debug, PartialEq, Identifiable, Associations, Queryable, Insertable)]
#[belongs_to(Person, foreign_key = "person_id")]
#[table_name = "post"]
pub struct Post {
    pub id: Uuid,
    pub person_id: Uuid,
    pub text: String,
    pub create_at: NaiveDateTime,
    pub update_at: NaiveDateTime,
}

#[derive(Debug, Queryable, Insertable)]
#[table_name = "post"]
pub struct NewPost {
    pub person_id: Uuid,
    pub text: String,
}

#[derive(Debug)]
pub struct PostWithPerson {
    person_id: uuid::Uuid,
    name: String,
    post_id: uuid::Uuid,
    text: String,
    create_at: NaiveDateTime,
    update_at: NaiveDateTime,
}

#[derive(Debug, Queryable)]
pub struct PostWithPersonTuple(Post, Person);

impl Person {
    pub fn new(name: String) -> Person {
        let now = Utc::now().naive_utc();

        Person {
            id: Uuid::new_v4(),
            name,
            create_at: now,
            update_at: now,
        }
    }
}

impl Post {
    pub fn new(person_id: Uuid, text: String) -> Post {
        let now = Utc::now().naive_utc();

        Post {
            id: Uuid::new_v4(),
            person_id,
            text,
            create_at: now,
            update_at: now,
        }
    }
}

impl PostWithPersonTuple {
    pub fn flatten(self) -> PostWithPerson {
        PostWithPerson {
            person_id: self.1.id,
            name: self.1.name,
            post_id: self.0.id,
            text: self.0.text,
            create_at: self.1.create_at,
            update_at: self.1.update_at,
        }
    }
}
