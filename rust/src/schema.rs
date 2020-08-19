table! {
    person (id) {
        id -> Uuid,
        name -> Varchar,
    }
}

table! {
    post (id) {
        id -> Uuid,
        person_id -> Uuid,
        text -> Varchar,
    }
}

joinable!(post -> person(person_id));
allow_tables_to_appear_in_same_query!(post, person);
