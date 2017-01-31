/*** Service to Display Properties ***/

//import
import { Injectable }     from '@angular/core';
import { Observable } from 'rxjs/Rx';
import { Subject } from 'rxjs/Subject';
import { Pod } from '../model/pod'

// Import RxJs required methods
import 'rxjs/add/operator/map';
import 'rxjs/add/operator/catch';

@Injectable()
export class DisplayPropertiesService {
     
     constructor () {}
     // private instance variable to hold list of properties and Observable
     private namespace: string;
     private subject: Subject<string> = new Subject<string>();

     setProperties(namespace: string): void {
         this.namespace = namespace;
         this.subject.next(namespace);
     }
     getProperties() : Observable<string> {
        return this.subject.asObservable();
     }
}