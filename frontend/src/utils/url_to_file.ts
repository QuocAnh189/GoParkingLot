export async function URLtoFile(url: string) {
  try {
    const response = await fetch(url);
    const blob = await response.blob();
    const file = new File([blob], url.split("/")[5].split("?")[0], {
      type: blob.type,
    });
    return file;
  } catch (error) {
    console.log(error);
  }
}
