// tslint:disable
/**
 * Tweetwatch Server
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * The version of the OpenAPI document: 0.0.1
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
/**
 * 
 * @export
 * @interface CreateStream
 */
export interface CreateStream {
    /**
     * 
     * @type {string}
     * @memberof CreateStream
     */
    track: string;
}

export function CreateStreamFromJSON(json: any): CreateStream {
    return CreateStreamFromJSONTyped(json, false);
}

export function CreateStreamFromJSONTyped(json: any, ignoreDiscriminator: boolean): CreateStream {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'track': json['track'],
    };
}

export function CreateStreamToJSON(value?: CreateStream | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'track': value.track,
    };
}


