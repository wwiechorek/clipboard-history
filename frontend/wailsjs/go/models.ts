export namespace main {
	
	export class Clip {
	    ID: number;
	    TSISO: string;
	    Content: string;
	
	    static createFrom(source: any = {}) {
	        return new Clip(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.ID = source["ID"];
	        this.TSISO = source["TSISO"];
	        this.Content = source["Content"];
	    }
	}

}

