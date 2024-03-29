export class User {
    id?: string;
    user_name: string;
    first_name: string;
    last_name: string;
    email: string;
    user_status: string;
    department: string;

    constructor(
        user_name: string,
        first_name: string,
        last_name: string,
        email: string,
        user_status: string,
        department: string
    ) {
        this.user_name = user_name;
        this.first_name = first_name;
        this.last_name = last_name;
        this.email = email;
        this.user_status = user_status;
        this.department = department;
    }
}
