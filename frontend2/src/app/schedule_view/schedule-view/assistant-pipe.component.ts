import { Pipe, PipeTransform } from '@angular/core';

@Pipe({
    name: 'assistantPipe',
    pure: false
})
export class AssistantPipe implements PipeTransform {
    transform(items: any[], assistantType: Object): any {
        if (!items || !assistantType) {
            return items;
        }
        // filter items array, items which match and return true will be
        // kept, false will be filtered out
        return items.filter((a) => a.type == assistantType);
    }
}