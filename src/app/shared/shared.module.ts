import { FontAwesomeModule } from '@fortawesome/angular-fontawesome';
import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { RouterModule } from '@angular/router';
import { MomentModule } from 'ngx-moment';
import { EditorModule } from '@tinymce/tinymce-angular';
import { ModalModule } from 'ngx-bootstrap';
import { FormsModule } from '@angular/forms';

import { MainLayoutComponent } from './layouts/main-layout/main-layout.component';
import { FullScreenHeroComponent } from './components/full-screen-hero/full-screen-hero.component';
import { TextSectionComponent } from './components/text-section/text-section.component';
import { InlineEditorComponent } from './components/inline-editor/inline-editor.component';
import { SafeHtmlPipe } from './pipes/safe-html.pipe';
import { DeletedPipe } from './pipes/deleted.pipe';

@NgModule({
  imports: [
    CommonModule,
    RouterModule,
    FontAwesomeModule,
    EditorModule,
    ModalModule.forRoot(),
    FormsModule

  ],
  declarations: [MainLayoutComponent, FullScreenHeroComponent, TextSectionComponent, InlineEditorComponent, SafeHtmlPipe, DeletedPipe],
  exports: [MainLayoutComponent, FullScreenHeroComponent,TextSectionComponent,DeletedPipe]
})
export class SharedModule { }
