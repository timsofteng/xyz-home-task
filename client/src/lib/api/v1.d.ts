/**
 * This file was auto-generated by openapi-typescript.
 * Do not make direct changes to the file.
 */

export interface paths {
    "/books": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        /** Get all books */
        get: operations["getBooks"];
        put?: never;
        post?: never;
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
}
export type webhooks = Record<string, never>;
export interface components {
    schemas: {
        ErrorResponse: {
            /**
             * @description Description of the error
             * @example Internal server error occurred
             */
            message: string;
        };
        BookBase: {
            /**
             * @description Unique ID
             * @example dIewflkjW
             */
            id: string;
            /**
             * @description Title of the book
             * @example Vim for everyone
             */
            title: string;
            /**
             * @description Description of the book
             * @example Blah blah blah and so on
             */
            description: string;
            /**
             * @description Pages in the book
             * @example 245
             */
            pages: number;
            /**
             * Format: float
             * @description Price of the book
             * @example 30
             */
            price: number;
            /** @description Currency of price */
            currency: string;
            /** @description Thumbnail URL */
            thumbnail: string;
        };
        BooksListResponse: {
            items: components["schemas"]["BookBase"][];
        };
    };
    responses: never;
    parameters: never;
    requestBodies: never;
    headers: never;
    pathItems: never;
}
export type $defs = Record<string, never>;
export interface operations {
    getBooks: {
        parameters: {
            query: {
                /** @description Filter query */
                q: string;
            };
            header?: never;
            path?: never;
            cookie?: never;
        };
        requestBody?: never;
        responses: {
            /** @description success */
            200: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["BooksListResponse"];
                };
            };
            /** @description Internal server error */
            500: {
                headers: {
                    [name: string]: unknown;
                };
                content: {
                    "application/json": components["schemas"]["ErrorResponse"];
                };
            };
        };
    };
}
