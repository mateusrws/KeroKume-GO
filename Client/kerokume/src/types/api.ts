export type ApiListResponse<T> = {
  message: string
  data?: T[]
}

export type ApiSingleResponse<T> = {
  message: string
  data?: T
}
