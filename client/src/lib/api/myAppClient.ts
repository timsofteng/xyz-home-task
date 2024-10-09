import createFetchClient from "openapi-fetch";
import createClient from "openapi-react-query";
import type { paths } from "./v1";
import { SERVER_BASE_URL } from "../../config";

const fetchClient = createFetchClient<paths>({
  baseUrl: SERVER_BASE_URL,
});

export const $api = createClient(fetchClient);
