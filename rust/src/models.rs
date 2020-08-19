use crate::schema::person;
use crate::schema::post;
use uuid::Uuid;

#[derive(Debug, Identifiable, Queryable, Insertable, PartialEq)]
#[table_name = "person"]
pub struct Person {
    pub id: Uuid,
    pub name: String,
}

#[derive(Debug, Queryable, Insertable)]
#[table_name = "person"]
pub struct NewPerson {
    pub name: String,
}

#[derive(Debug, Identifiable, Associations, Queryable, Insertable, PartialEq)]
#[belongs_to(Person, foreign_key = "person_id")]
#[table_name = "post"]
pub struct Post {
    pub id: Uuid,
    pub person_id: Uuid,
    pub text: String,
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
}

#[derive(Debug, Queryable)]
pub struct PostWithPersonTuple(Post, Person);

impl Person {
    pub fn new(name: String) -> Person {
        Person {
            id: Uuid::new_v4(),
            name,
        }
    }
}

impl Post {
    pub fn new(person_id: Uuid, text: String) -> Post {
        Post {
            id: Uuid::new_v4(),
            person_id,
            text,
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
        }
    }
}
