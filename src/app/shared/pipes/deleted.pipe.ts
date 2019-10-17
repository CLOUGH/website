import { Pipe, PipeTransform } from '@angular/core';

@Pipe({
  name: 'deleted'
})
export class DeletedPipe implements PipeTransform {

  transform(items: any[], ...args: any[]): any {
    if (!items) {
      return [];

    }
    return items.filter(item => !item.deleted);
  }

}
