import { CommonModule } from '@angular/common';
import { Component, OnInit } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { Note } from '../../services/note';

@Component({
  selector: 'app-notes',
  imports: [CommonModule, FormsModule],
  templateUrl: './notes.html',
  styleUrl: './notes.css',
})
export class Notes implements OnInit {

  notes: Note[] = [];
  newNote = '';

  constructor(private noteService: Note) {}

  ngOnInit(): void {
    this.loadNotes();
  }

  loadNotes(): void {
    this.noteService.getNotes().subscribe((data: Note[]) => {
      this.notes = data;
    });
  }

  addNote(): void {
    if (!this.newNote.trim()) return;

    this.noteService.createNote(this.newNote).subscribe(() => {
      this.newNote = '';
      this.loadNotes();
    });
  }

  deleteNote(id: number): void {
    this.noteService.deleteNote(id).subscribe(() => {
      this.loadNotes();
    });
  }
}
