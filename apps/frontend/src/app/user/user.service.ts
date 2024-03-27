import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Observable } from 'rxjs';
import {User} from './user'

@Injectable({
  providedIn: 'root'
})
export class UserService {
  private hostUrl=`http://localhost/api`;

  constructor(private http: HttpClient) {}

  getUsers(search: string, limit: number, page: number): Observable<any> {
    return this.http.get(`${this.hostUrl}/users?search=${search}&limit=${limit}&page=${page}`)
  }

  getUser(id: number): Observable<any>{
    return this.http.get(`${this.hostUrl}/user/${id}`)
  }

  createUser(user: User): Observable<any>{
    return this.http.post(`${this.hostUrl}/user`, user)
  }

  updateUser(id: number, user: User): Observable<any>{
    return this.http.put(`${this.hostUrl}/user`, user)
  }

  deleteUser(id: number): Observable<any>{
    return this.http.delete(`${this.hostUrl}/user/${id}`)
  }
}
