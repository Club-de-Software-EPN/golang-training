type Team = "CS" | "AI" | "WEB"

class ClubMember {
    readonly _id: number;
    private _fullname: string;
    private _email: string;
    private _team: Team;

    constructor(id: number, fullname: string, email: string, team: Team){
        this._id = id;
        this._fullname = fullname;
        this._email = email;
        this._team = team;
    }

    get id(){
        return this._id;
    }

    get fullname(){
        return this._fullname;
    }

    get email(){
        return this._email;
    }

    get team(){
        return this._team;
    }

    set fullname(name: string){
        this._fullname = name;
    }

    set email(email: string){
        this._email = email;
    }

    set team(team: Team){
        this._team = team;
    }
}

// Create an object
const securityMember = new ClubMember(0, "Alejandro Llanganate", "example@gmail.com", "CS");
console.log(securityMember)

// Use getters
console.log("get email:", securityMember.email)

// Use setters
securityMember.email = "alelejandro@gmail.com"
console.log("get new email:", securityMember.email)

// readonly
// securityMember.id = 5

