import { FontAwesomeModule } from '@fortawesome/angular-fontawesome';
import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { RouterModule } from '@angular/router';
import { MomentModule } from 'ngx-moment';
import { CKEditorModule } from '@ckeditor/ckeditor5-angular';
import { QuillModule } from 'ngx-quill';
import { EditorModule } from '@tinymce/tinymce-angular';

import { MainLayoutComponent } from './layouts/main-layout/main-layout.component';
import { FullScreenHeroComponent } from './components/full-screen-hero/full-screen-hero.component';

@NgModule({
  imports: [
    CommonModule,
    RouterModule,
    FontAwesomeModule,
    CKEditorModule,
    EditorModule,
    QuillModule.forRoot({
      modules: {
        syntax: true,
      }
    })
  ],
  declarations: [MainLayoutComponent, FullScreenHeroComponent],
  exports: [MainLayoutComponent, FullScreenHeroComponent]
})
export class SharedModule { }
