import { Component, OnInit, Input } from '@angular/core';
import { post } from 'selenium-webdriver/http';
import { Post } from '../../models/post';

@Component({
  selector: 'app-post-preview',
  templateUrl: './post-preview.component.html',
  styleUrls: ['./post-preview.component.scss']
})
export class PostPreviewComponent implements OnInit {

  @Input() post: Post;

  constructor() { }

  ngOnInit() {
  }

}
