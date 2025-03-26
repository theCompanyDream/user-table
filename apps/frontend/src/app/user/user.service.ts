import { Injectable } from '@angular/core';
import { environment } from '../../environments';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import {User} from './user'

@Injectable({
  providedIn: 'root'
})
export class UserService {
  private hostUrl= environment.apiUrl;

  constructor(private http: HttpClient) {}

  getUsers(search: string, limit: number, page: number): Observable<any> {
    return this.http.get<User>(`${this.hostUrl}/users?search=${search}&limit=${limit}&page=${page}`)
  }

  getUser(id: string): Observable<any>{
    return this.http.get(`${this.hostUrl}/user/${id}`)
  }

  createUser(user: User): Observable<any>{
    return this.http.post(`${this.hostUrl}/user`, user)
  }

  updateUser(id: string, user: User): Observable<any>{
    return this.http.put(`${this.hostUrl}/user/${id}`, user)
  }

  deleteUser(id: string): Observable<any>{
    return this.http.delete(`${this.hostUrl}/user/${id}`)
  }
}
