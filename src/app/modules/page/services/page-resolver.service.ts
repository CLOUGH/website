import { Injectable } from '@angular/core';
import { Location } from '@angular/common';
import { Resolve, ActivatedRouteSnapshot, RouterStateSnapshot, Router } from '@angular/router';
import { of } from 'rxjs';
import { map } from 'rxjs/operators';
import { PageService } from './page.service';

@Injectable({
  providedIn: 'root'
})
export class PageResolverService implements Resolve<any> {

  constructor(private pageService: PageService, private router: Router, private location: Location) { }

  resolve(route: ActivatedRouteSnapshot, state: RouterStateSnapshot) {
    const path = route.url.join('/');
    console.log({path});
    return this.pageService.findPage(path).pipe(map(page => {
      if (!page) {
        // this.location.replaceState(path);

        this.router.navigateByUrl('error/404',{ skipLocationChange: false });
        return;
      }
      return page;
    }));
  }
}
