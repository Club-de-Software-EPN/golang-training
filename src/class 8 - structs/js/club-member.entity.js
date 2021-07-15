"use strict";
class ClubMember {
    constructor(id, fullname, email, team) {
        this._id = id;
        this._fullname = fullname;
        this._email = email;
        this._team = team;
    }
    get id() {
        return this._id;
    }
    get fullname() {
        return this._fullname;
    }
    get email() {
        return this._email;
    }
    get team() {
        return this._team;
    }
    set fullname(name) {
        this._fullname = name;
    }
    set email(email) {
        this._email = email;
    }
    set team(team) {
        this._team = team;
    }
}
// Create an object
const securityMember = new ClubMember(0, "Alejandro Llanganate", "example@gmail.com", "CS");
console.log(securityMember);
// Use getters
console.log("get email:", securityMember.email);
// Use setters
securityMember.email = "alelejandro@gmail.com";
console.log("get new email:", securityMember.email);
try {
    // idquito/en/home-en/
    securityMember.id = 5;
}
catch (_a) {
    console.log("sotty");
}
