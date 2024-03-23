import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Observable } from 'rxjs';
import {User} from './user'

@Injectable({
  providedIn: 'root'
})
export class UserService {
  private hostUrl=`${process.env['BACKEND_NAME']}:${process.env['BACKEND_PORT']}`;

  constructor(private http: HttpClient) {}

  getUsers(search: string, limit: BigInteger, page: BigInteger): Observable<any> {
    return this.http.get(`${this.hostUrl}/users?search=${search}&limit=${limit}&page=${page}`)
  }

  getUser(id: BigInteger): Observable<any>{
    return this.http.get(`${this.hostUrl}/user/${id}`)
  }

  createUser(user: User): Observable<User>{
    return this.http.post(`${this.hostUrl}/user`, user)
  }

  updateUser(id: BigInteger, user: User): Observable<User>{
    return this.http.put(`${this.hostUrl}/user`, user)
  }

  deleteUser(id: BigInteger): Observable<any>{
    return this.http.delete(`${this.hostUrl}/user/${id}`)
  }

}
