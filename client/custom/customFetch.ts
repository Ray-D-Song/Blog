const customFetch = async (_fetch: any) => {
  const res = await _fetch
  return res.data.value
}

export default customFetch