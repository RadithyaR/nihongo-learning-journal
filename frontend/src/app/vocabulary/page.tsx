"use client"

import { useCallback, useEffect, useState } from "react"
import { api } from "@/lib/axios"
import { Plus, Search, Edit, Trash2, Heart, Loader2, Filter, ArrowUpDown } from "lucide-react"
import { TagPicker } from "@/components/tag-picker"

import { Button } from "@/components/ui/button"
import { Input } from "@/components/ui/input"
import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@/components/ui/table"
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from "@/components/ui/dialog"
import { Label } from "@/components/ui/label"
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select"

interface TagItem {
  id: string
  name: string
  color?: string | null
}

interface Meaning {
  id: string
  meaning: string
  order_number: number
}

interface Vocabulary {
  id: string
  word: string
  reading?: string
  source?: string
  note?: string
  status?: string
  favourite: boolean
  meanings: Meaning[]
  tags?: TagItem[]
}

export default function VocabularyPage() {
  const [vocabularies, setVocabularies] = useState<Vocabulary[]>([])
  const [loading, setLoading] = useState(true)
  const [search, setSearch] = useState("")
  const [filterFavourite, setFilterFavourite] = useState<string>("all")
  const [filterTag, setFilterTag] = useState<string>("all")
  const [filterSort, setFilterSort] = useState<string>("newest")
  const [tags, setTags] = useState<TagItem[]>([])

  const [isDialogOpen, setIsDialogOpen] = useState(false)
  const [editId, setEditId] = useState<string | null>(null)
  const [formData, setFormData] = useState({
    word: "",
    reading: "",
    meanings: [""],
    source: "",
    note: ""
  })
  const [isSubmitting, setIsSubmitting] = useState(false)

  const fetchTags = async () => {
    try {
      const response = await api.get("/tags")
      setTags(response.data.data)
    } catch (err) {
      console.error("Failed to load tags", err)
    }
  }

  const fetchVocabularies = useCallback(async () => {
    setLoading(true)
    try {
      const params = new URLSearchParams()
      if (search) params.set("search", search)
      if (filterFavourite === "true") params.set("favourite", "true")
      if (filterTag !== "all") params.set("tagId", filterTag)
      if (filterSort) params.set("sort", filterSort)

      const queryString = params.toString()
      const response = await api.get(`/vocabularies${queryString ? `?${queryString}` : ""}`)
      setVocabularies(response.data.data)
    } catch (err) {
      console.error("Failed to load vocabularies", err)
    } finally {
      setLoading(false)
    }
  }, [search, filterFavourite, filterTag, filterSort])

  useEffect(() => {
    fetchTags()
  }, [])

  useEffect(() => {
    const delayDebounceFn = setTimeout(() => {
      fetchVocabularies()
    }, 500)
    return () => clearTimeout(delayDebounceFn)
  }, [fetchVocabularies])

  const toggleFavourite = async (id: string) => {
    try {
      await api.patch(`/vocabularies/${id}/favourite`)
      setVocabularies(vocabularies.map(v => 
        v.id === id ? { ...v, favourite: !v.favourite } : v
      ))
    } catch (err) {
      console.error("Failed to toggle favourite", err)
    }
  }

  const deleteVocabulary = async (id: string) => {
    if (!confirm("Are you sure you want to delete this item?")) return
    
    try {
      await api.delete(`/vocabularies/${id}`)
      setVocabularies(vocabularies.filter(v => v.id !== id))
    } catch (err) {
      console.error("Failed to delete", err)
    }
  }

  const handleMeaningChange = (index: number, value: string) => {
    const newMeanings = [...formData.meanings]
    newMeanings[index] = value
    setFormData({ ...formData, meanings: newMeanings })
  }

  const addMeaningField = () => {
    setFormData({ ...formData, meanings: [...formData.meanings, ""] })
  }

  const removeMeaningField = (index: number) => {
    if (formData.meanings.length <= 1) return
    const newMeanings = [...formData.meanings]
    newMeanings.splice(index, 1)
    setFormData({ ...formData, meanings: newMeanings })
  }

  const openEdit = (vocab: Vocabulary) => {
    setFormData({
      word: vocab.word,
      reading: vocab.reading || "",
      meanings: vocab.meanings && vocab.meanings.length > 0 
        ? vocab.meanings.map(m => m.meaning) 
        : [""],
      source: vocab.source || "",
      note: vocab.note || ""
    })
    setEditId(vocab.id)
    setIsDialogOpen(true)
  }

  const handleOpenChange = (open: boolean) => {
    setIsDialogOpen(open)
    if (!open) {
      setEditId(null)
      setFormData({ word: "", reading: "", meanings: [""], source: "", note: "" })
    }
  }

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault()
    setIsSubmitting(true)
    try {
      const filteredMeanings = formData.meanings.filter(m => m.trim() !== "")
      const payload = {
        ...formData,
        meanings: filteredMeanings.length > 0 ? filteredMeanings : ["Undefined"]
      }
      
      if (editId) {
        await api.put(`/vocabularies/${editId}`, payload)
      } else {
        await api.post("/vocabularies", payload)
      }
      setIsDialogOpen(false)
      setEditId(null)
      setFormData({ word: "", reading: "", meanings: [""], source: "", note: "" })
      fetchVocabularies()
    } catch (err) {
      console.error("Failed to save", err)
    } finally {
      setIsSubmitting(false)
    }
  }

  return (
    <div className="container max-w-6xl mx-auto py-8 px-4 space-y-6">
      <div className="flex flex-col md:flex-row justify-between items-start md:items-center gap-4">
        <div>
          <h1 className="text-3xl font-bold tracking-tight">Vocabulary</h1>
          <p className="text-muted-foreground">Manage your Japanese vocabulary collection.</p>
        </div>
        
        <Dialog open={isDialogOpen} onOpenChange={handleOpenChange}>
          <DialogTrigger asChild>
            <Button>
              <Plus className="mr-2 h-4 w-4" /> Add New Word
            </Button>
          </DialogTrigger>
          <DialogContent className="sm:max-w-[500px]">
            <DialogHeader>
              <DialogTitle>{editId ? "Edit Vocabulary" : "Add Vocabulary"}</DialogTitle>
              <DialogDescription>
                {editId ? "Edit your Japanese word details." : "Add a new Japanese word to your study collection."}
              </DialogDescription>
            </DialogHeader>
            <form onSubmit={handleSubmit} className="space-y-4 py-4">
              <div className="space-y-2">
                <Label htmlFor="word">Word (Kanji/Kana) <span className="text-destructive">*</span></Label>
                <Input 
                  id="word" 
                  value={formData.word} 
                  onChange={(e) => setFormData({...formData, word: e.target.value})} 
                  required 
                />
              </div>
              <div className="space-y-2">
                <Label htmlFor="reading">Reading (Furigana/Kana)</Label>
                <Input 
                  id="reading" 
                  value={formData.reading} 
                  onChange={(e) => setFormData({...formData, reading: e.target.value})} 
                />
              </div>
              <div className="space-y-2">
                <Label>Meanings <span className="text-destructive">*</span></Label>
                {formData.meanings.map((meaning, index) => (
                  <div key={index} className="flex gap-2 mb-2">
                    <Input 
                      value={meaning} 
                      onChange={(e) => handleMeaningChange(index, e.target.value)}
                      placeholder={`Meaning ${index + 1}`}
                      required={index === 0}
                    />
                    {formData.meanings.length > 1 && (
                      <Button type="button" variant="ghost" size="icon" onClick={() => removeMeaningField(index)}>
                        <Trash2 className="h-4 w-4 text-destructive" />
                      </Button>
                    )}
                  </div>
                ))}
                <Button type="button" variant="outline" size="sm" onClick={addMeaningField} className="mt-2">
                  <Plus className="mr-2 h-3 w-3" /> Add Meaning
                </Button>
              </div>
              <div className="space-y-2">
                <Label htmlFor="source">Source (Optional)</Label>
                <Input 
                  id="source" 
                  value={formData.source} 
                  onChange={(e) => setFormData({...formData, source: e.target.value})} 
                  placeholder="e.g. Genki I, Anime, etc."
                />
              </div>
              <div className="space-y-2">
                <Label htmlFor="note">Notes (Optional)</Label>
                <Input 
                  id="note" 
                  value={formData.note} 
                  onChange={(e) => setFormData({...formData, note: e.target.value})} 
                />
              </div>
              <DialogFooter>
                <Button type="submit" disabled={isSubmitting}>
                  {isSubmitting && <Loader2 className="mr-2 h-4 w-4 animate-spin" />}
                  Save Item
                </Button>
              </DialogFooter>
            </form>
          </DialogContent>
        </Dialog>
      </div>

      <div className="flex items-center relative max-w-sm">
        <Search className="absolute left-3 top-1/2 -translate-y-1/2 h-4 w-4 text-muted-foreground" />
        <Input 
          placeholder="Search vocabulary..." 
          className="pl-9"
          value={search}
          onChange={(e) => setSearch(e.target.value)}
        />
      </div>

      <div className="flex flex-wrap items-center gap-3">
        <div className="flex items-center gap-2">
          <Filter className="h-4 w-4 text-muted-foreground" />
          <span className="text-sm text-muted-foreground">Filter:</span>
        </div>

        <Select value={filterFavourite} onValueChange={setFilterFavourite}>
          <SelectTrigger className="w-[140px]">
            <SelectValue placeholder="Favourite" />
          </SelectTrigger>
          <SelectContent>
            <SelectItem value="all">All Items</SelectItem>
            <SelectItem value="true">Favourited</SelectItem>
          </SelectContent>
        </Select>

        <Select value={filterTag} onValueChange={setFilterTag}>
          <SelectTrigger className="w-[140px]">
            <SelectValue placeholder="Tag" />
          </SelectTrigger>
          <SelectContent>
            <SelectItem value="all">All Tags</SelectItem>
            {tags.map((tag) => (
              <SelectItem key={tag.id} value={tag.id}>
                {tag.name}
              </SelectItem>
            ))}
          </SelectContent>
        </Select>

        <div className="flex items-center gap-2">
          <ArrowUpDown className="h-4 w-4 text-muted-foreground" />
          <Select value={filterSort} onValueChange={setFilterSort}>
            <SelectTrigger className="w-[140px]">
              <SelectValue placeholder="Sort" />
            </SelectTrigger>
            <SelectContent>
              <SelectItem value="newest">Newest First</SelectItem>
              <SelectItem value="oldest">Oldest First</SelectItem>
            </SelectContent>
          </Select>
        </div>
      </div>

      <div className="border rounded-md">
        <Table>
          <TableHeader>
            <TableRow>
              <TableHead className="w-[50px]"></TableHead>
              <TableHead>Word</TableHead>
              <TableHead>Reading</TableHead>
              <TableHead>Meanings</TableHead>
              <TableHead>Status</TableHead>
              <TableHead>Tags</TableHead>
              <TableHead className="text-right">Actions</TableHead>
            </TableRow>
          </TableHeader>
          <TableBody>
            {loading ? (
              <TableRow>
                <TableCell colSpan={7} className="text-center py-10 text-muted-foreground">
                  <Loader2 className="h-6 w-6 animate-spin mx-auto" />
                </TableCell>
              </TableRow>
            ) : vocabularies.length === 0 ? (
              <TableRow>
                <TableCell colSpan={7} className="text-center py-10 text-muted-foreground">
                  No vocabulary found. Add your first word!
                </TableCell>
              </TableRow>
            ) : (
              vocabularies.map((item) => (
                <TableRow key={item.id}>
                  <TableCell>
                    <button 
                      onClick={() => toggleFavourite(item.id)}
                      className={`hover:bg-muted p-1.5 rounded-full transition-colors ${item.favourite ? 'text-red-500' : 'text-muted-foreground'}`}
                    >
                      <Heart className="h-4 w-4" fill={item.favourite ? "currentColor" : "none"} />
                    </button>
                  </TableCell>
                  <TableCell className="font-bold text-lg">{item.word}</TableCell>
                  <TableCell className="text-muted-foreground">{item.reading || "-"}</TableCell>
                  <TableCell>
                    {item.meanings?.map(m => m.meaning).join(", ") || "-"}
                  </TableCell>
                  <TableCell>
                    <span className="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium bg-primary/10 text-primary">
                      {item.status || "NEW"}
                    </span>
                  </TableCell>
                  <TableCell>
                    <TagPicker
                      itemId={item.id}
                      itemType="VOCABULARY"
                      attachedTags={item.tags}
                      onTagsChange={(tags) => {
                        setVocabularies(vocabularies.map(v => v.id === item.id ? { ...v, tags } : v))
                      }}
                    />
                  </TableCell>
                  <TableCell className="text-right">
                    <Button variant="ghost" size="icon" title="Edit" onClick={() => openEdit(item)}>
                      <Edit className="h-4 w-4 text-muted-foreground" />
                    </Button>
                    <Button variant="ghost" size="icon" title="Delete" onClick={() => deleteVocabulary(item.id)}>
                      <Trash2 className="h-4 w-4 text-destructive" />
                    </Button>
                  </TableCell>
                </TableRow>
              ))
            )}
          </TableBody>
        </Table>
      </div>
    </div>
  )
}
