
import type { BlogPost, Category, Tag, MediaFile } from '@/types';

export const BLOG_TITLE = "Voocel.";
export const BLOG_SUBTITLE = "Notes from the Aether";
export const AUTHOR_NAME = "Voocel";

export const HERO_CONTENT = {
  title: 'The Aesthetics of Code',
  excerpt: 'Exploring the silent rhythm of algorithms and why the most beautiful functions are often the most dangerous.',
  cover: 'https://images.unsplash.com/photo-1541963463532-d68292c34b19?q=80&w=2576&auto=format&fit=crop',
};

export const MOCK_CATEGORIES: Category[] = [
  { id: 1, name: 'Philosophy', slug: 'philosophy', count: 3 },
  { id: 2, name: 'Design', slug: 'design', count: 3 },
  { id: 3, name: 'Engineering', slug: 'engineering', count: 1 },
  { id: 4, name: 'Travel', slug: 'travel', count: 1 },
];

export const MOCK_TAGS: Tag[] = [
  { id: 1, name: 'AI' },
  { id: 2, name: 'Minimalism' },
  { id: 3, name: 'React' },
  { id: 4, name: 'Tokyo' },
  { id: 5, name: 'Cyberpunk' },
  { id: 6, name: 'Typography' },
  { id: 7, name: 'Web' },
];

export const MOCK_FILES: MediaFile[] = [
  { id: 1, url: 'https://images.unsplash.com/photo-1541963463532-d68292c34b19?q=80&w=2576&auto=format&fit=crop', name: 'Golden Sands', type: 'image', date: '2023-10-01' },
  { id: 2, url: 'https://images.unsplash.com/photo-1493770348161-369560ae357d?q=80&w=2670&auto=format&fit=crop', name: 'Minimal Shadow', type: 'image', date: '2023-11-12' },
  { id: 3, url: 'https://images.unsplash.com/photo-1540959733332-eab4deabeeaf?q=80&w=2694&auto=format&fit=crop', name: 'Warm Tokyo', type: 'image', date: '2023-12-05' },
  { id: 4, url: 'https://images.unsplash.com/photo-1487958449943-2429e8be8625?q=80&w=2670&auto=format&fit=crop', name: 'Concrete Structure', type: 'image', date: '2024-01-15' },
  { id: 5, url: 'https://images.unsplash.com/photo-1506784983877-45594efa4cbe?q=80&w=2668&auto=format&fit=crop', name: 'Nature Detail', type: 'image', date: '2024-02-10' },
  { id: 6, url: 'https://images.unsplash.com/photo-1618005182384-a83a8bd57fbe?auto=format&fit=crop&w=2564&q=80', name: 'Abstract Liquid', type: 'image', date: '2024-03-01' },
  { id: 7, url: 'https://images.unsplash.com/photo-1461360370896-922624d12aa1?auto=format&fit=crop&w=2000&q=80', name: 'Old Books', type: 'image', date: '2024-03-10' },
  { id: 8, url: 'https://images.unsplash.com/photo-1505682634904-d7c8d95cdc50?auto=format&fit=crop&w=2000&q=80', name: 'Digital Rain', type: 'image', date: '2024-03-15' },
];

export const MOCK_POSTS: BlogPost[] = [
  {
    id: 1,
    slug: 'the-aesthetics-of-code',
    title: 'The Aesthetics of Code',
    excerpt: 'Exploring the silent rhythm of algorithms and why the most beautiful functions are often the most dangerous.',
    content: `> *Generative art has moved from the fringes to the center of the design world.*

Generative art has moved from the fringes to the center of the design world. But what does it mean for code to be beautiful?

## The Algorithm as Artist

When we write code, we usually think about utility. Does it work? Is it fast? But there is a rhythm to a well-written function.

\`\`\`javascript
const beauty = (code) => {
  return code.isClean() && code.isEfficient();
}
\`\`\`

True beauty lies in **simplicity**. The most elegant solutions are often the ones that do less, not more.

### The Void
We often fear the empty space in our editors, rushing to fill it with logic. But in that void, there is potential.`,
    author: AUTHOR_NAME,
    publishAt: '2024-05-15T09:00:00+08:00',
    categoryId: 1,
    category: 'Philosophy',
    readTime: '5 min read',
    cover: MOCK_FILES[0].url,
    tags: ['AI', 'Minimalism'],
    views: 1240,
    status: 'published'
  },
  {
    id: 2,
    slug: 'silence-in-the-signal',
    title: 'Silence in the Signal',
    excerpt: 'Why minimalism is not just an aesthetic choice, but a survival mechanism in the age of infinite information.',
    content: `> *We are drowning in notifications.*

We are drowning in notifications, feeds, and constant updates. Minimalism is not just about white space on a page; it is about white space in the mind.

## Design as Subtraction

To design is to remove.

- Remove the noise
- Remove the friction
- Remove the ego

When we strip away the non-essential, we are left with the truth.`,
    author: AUTHOR_NAME,
    publishAt: '2024-05-10T09:00:00+08:00',
    categoryId: 2,
    category: 'Design',
    readTime: '4 min read',
    cover: MOCK_FILES[1].url,
    tags: ['Minimalism'],
    views: 856,
    status: 'published'
  },
  {
    id: 3,
    slug: 'tokyo-golden-hour',
    title: 'Tokyo: Golden Hour',
    excerpt: 'A visual journey through the warm, lantern-lit streets and hidden alleyways of the metropolis.',
    content: `> *Tokyo is a sensory overload, but at golden hour, it breathes.*

The light hits the Shibuya crossing in a way that turns concrete into gold.

## The Light

It wasn't just the sun; it was the reflection off millions of glass panels, creating a diffuse glow that enveloped the city.

![Tokyo Street](${MOCK_FILES[2].url})

I found myself walking for hours, camera in hand, looking for the quiet moments between the chaos.`,
    author: AUTHOR_NAME,
    publishAt: '2024-04-22T09:00:00+08:00',
    categoryId: 4,
    category: 'Travel',
    readTime: '6 min read',
    cover: MOCK_FILES[2].url,
    tags: ['Tokyo'],
    views: 2301,
    status: 'published'
  },
  {
    id: 4,
    slug: 'the-architecture-of-data',
    title: 'The Architecture of Data',
    excerpt: 'Breaking down the hype and understanding the architectural shift in the modern web ecosystem.',
    content: `> *React is evolving, and so are we.*

The shift to server components represents a fundamental change in how we think about the boundary between client and server.

## The Waterfall

We used to fear the waterfall. Now, we orchestrate it.

\`\`\`tsx
async function Page() {
  const data = await fetchData();
  return <Component data={data} />;
}
\`\`\`

This is not just a syntax change; it is a mental model shift.`,
    author: AUTHOR_NAME,
    publishAt: '2024-03-15T09:00:00+08:00',
    categoryId: 3,
    category: 'Engineering',
    readTime: '8 min read',
    cover: MOCK_FILES[3].url,
    tags: ['React'],
    views: 1540,
    status: 'published'
  },
  {
    id: 5,
    slug: 'the-analog-web',
    title: 'The Analog Web',
    excerpt: 'Revisiting the textures of early computing and how to bring tactile feeling back to digital interfaces.',
    content: `> *Before flat design took over, the web had texture.*

We need to bring back the feeling of material. Not skeuomorphism, but **soul**.

Buttons should feel clickable. Transitions should feel physical. The web is a medium, just like paper or clay.`,
    author: AUTHOR_NAME,
    publishAt: '2024-03-01T09:00:00+08:00',
    categoryId: 2,
    category: 'Design',
    readTime: '5 min read',
    cover: MOCK_FILES[5].url,
    tags: [],
    views: 920,
    status: 'published'
  },
  {
    id: 7,
    slug: 'digital-gardens',
    title: 'Digital Gardens',
    excerpt: 'Why we should stop building "streams" and start cultivating "gardens" of knowledge that grow over time.',
    content: `> *The stream is ephemeral. The garden is eternal.*

We spend too much time feeding the feed. We should be tending to our own digital gardens.

A garden is:
1. Slower
2. More thoughtful
3. Interconnected

Let's build spaces that age well.`,
    author: AUTHOR_NAME,
    publishAt: '2024-02-28T09:00:00+08:00',
    categoryId: 1,
    category: 'Philosophy',
    readTime: '6 min read',
    cover: MOCK_FILES[4].url,
    tags: ['Web'],
    views: 450,
    status: 'published'
  },
  {
    id: 8,
    slug: 'typography-as-voice',
    title: 'Typography as Voice',
    excerpt: 'How typefaces subconsciously alter the meaning of the words we read, and the responsibility of the designer.',
    content: `> *Type is the clothes words wear.*

You wouldn't wear a tuxedo to a beach party. Why use Helvetica for a romance novel?

**Serifs** carry history. **Sans-serifs** carry modernity. Choose wisely.`,
    author: AUTHOR_NAME,
    publishAt: '2024-02-15T09:00:00+08:00',
    categoryId: 2,
    category: 'Design',
    readTime: '4 min read',
    cover: MOCK_FILES[6].url,
    tags: ['Typography'],
    views: 1102,
    status: 'published'
  },
  {
    id: 9,
    slug: 'the-slow-web',
    title: 'The Slow Web',
    excerpt: 'A manifesto for building websites that respect attention, bandwidth, and the user\'s cognitive load.',
    content: `> *Speed is a feature, but stillness is a virtue.*

We optimize for TTI (Time to Interactive). We should optimize for TTC (Time to Contemplation).

Make it fast, yes. But make it calm.`,
    author: AUTHOR_NAME,
    publishAt: '2024-02-01T09:00:00+08:00',
    categoryId: 1,
    category: 'Philosophy',
    readTime: '5 min read',
    cover: MOCK_FILES[7].url,
    tags: ['Minimalism', 'Web'],
    views: 890,
    status: 'published'
  },
  {
    id: 6,
    slug: 'notes-on-solitude',
    title: 'Notes on Solitude',
    excerpt: 'A personal reflection on working alone and the silence required to build great things.',
    content: '# Notes on Solitude\n\nThis is a draft post about the importance of being alone with your thoughts...',
    author: AUTHOR_NAME,
    publishAt: '2024-02-10T09:00:00+08:00',
    categoryId: 1,
    category: 'Philosophy',
    readTime: '3 min read',
    cover: MOCK_FILES[4].url,
    tags: ['Minimalism'],
    views: 40,
    status: 'draft'
  }
];
