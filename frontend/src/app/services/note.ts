import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { environment } from '../../environments/environment.prod';

export interface Note {
  id: number;
  title: string;
}

@Injectable({
  providedIn: 'root'
})

export class Note {

  private baseUrl = environment.apiUrl;

  constructor(private http: HttpClient) {}

  getNotes(): Observable<Note[]> {
    return this.http.get<Note[]>(`${this.baseUrl}/notes`);
  }

  createNote(title: string): Observable<Note> {
    return this.http.post<Note>(`${this.baseUrl}/notes`, { title });
  }

  deleteNote(id: number): Observable<void> {
    return this.http.delete<void>(`${this.baseUrl}/notes/${id}`);
  }
}
