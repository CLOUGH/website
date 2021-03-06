import { FontAwesomeModule } from '@fortawesome/angular-fontawesome';
import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { RouterModule } from '@angular/router';
import { MomentModule } from 'ngx-moment';
import { EditorModule } from '@tinymce/tinymce-angular';
import { ModalModule, TimepickerModule } from 'ngx-bootstrap';
import { FormsModule } from '@angular/forms';

import { MainLayoutComponent } from './layouts/main-layout/main-layout.component';
import { FullScreenHeroComponent } from './components/full-screen-hero/full-screen-hero.component';
import { TextSectionComponent } from './components/text-section/text-section.component';
import { InlineEditorComponent } from './components/inline-editor/inline-editor.component';
import { SafeHtmlPipe } from './pipes/safe-html.pipe';
import { DeletedPipe } from './pipes/deleted.pipe';
import { VideoFullScreenHeroComponent } from './components/video-full-screen-hero/video-full-screen-hero.component';
import { TimePickerComponent } from './components/time-picker/time-picker.component';

@NgModule({
  imports: [
    CommonModule,
    RouterModule,
    FontAwesomeModule,
    EditorModule,
    ModalModule.forRoot(),
    FormsModule,    
    TimepickerModule.forRoot()

  ],
  declarations: [MainLayoutComponent,
    FullScreenHeroComponent,
    TextSectionComponent,
    InlineEditorComponent,
    SafeHtmlPipe,
    DeletedPipe,
    VideoFullScreenHeroComponent,
    TimePickerComponent
  ],
  exports: [
    MainLayoutComponent,
    FullScreenHeroComponent,
    TextSectionComponent,
    DeletedPipe,
    VideoFullScreenHeroComponent,
    TimePickerComponent
  ]
})
export class SharedModule { }
