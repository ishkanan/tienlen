let flashHandle: number | undefined = undefined

export function startFlashTitle(regular: string, alternate: string): void {
  if (flashHandle) {
    window.clearTimeout(flashHandle)
  }

  let n = 0
  const flasher = () => {
    flashHandle = undefined
    if (document.hasFocus()) {
      window.document.title = regular
      return
    }

    window.document.title = n === 0 ? regular : alternate
    n = (n + 1) % 2
    flashHandle = window.setTimeout(flasher, 1000)
  }

  flasher()
}

export function setTitle(title: string): void {
  if (flashHandle) {
    window.clearTimeout(flashHandle)
  }

  flashHandle = undefined
  window.document.title = title
}

export function ordinalise(n: number): string {
  const s = ['th', 'st', 'nd', 'rd']
  const v = n % 100
  return `${n}${s[(v - 20) % 10] || s[v] || s[0]}`
}
