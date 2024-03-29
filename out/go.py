
# python wrapper for package go within overall package hi
# This is what you import to use the package.
# File is generated by gopy. Do not edit.
# gopy build -vm=python3 -output=/out github.com/go-python/gopy/_examples/hi

# the following is required to enable dlopen to open the _go.so file
import os,sys,inspect,collections
cwd = os.getcwd()
currentdir = os.path.dirname(os.path.abspath(inspect.getfile(inspect.currentframe())))
os.chdir(currentdir)
import _hi
os.chdir(cwd)

# to use this code in your end-user python file, import it as follows:
# from hi import go
# and then refer to everything using go. prefix
# packages imported by this package listed below:


import collections
	
class GoClass(object):
	"""GoClass is the base class for all GoPy wrapper classes"""
	def __init__(self):
		self.handle = 0

# use go.nil for nil pointers 
nil = GoClass()

# need to explicitly initialize it
def main():
	global nil
	nil = GoClass()

main()

def Init():
	"""calls the GoPyInit function, which runs the 'main' code string that was passed using -main arg to gopy"""
	_hi.GoPyInit()

	


# ---- Types ---

# Python type for slice []bool
class Slice_bool(GoClass):
	""""""
	def __init__(self, *args, **kwargs):
		"""
		handle=A Go-side object is always initialized with an explicit handle=arg
		otherwise parameter is a python list that we copy from
		"""
		self.index = 0
		if len(kwargs) == 1 and 'handle' in kwargs:
			self.handle = kwargs['handle']
		elif len(args) == 1 and isinstance(args[0], GoClass):
			self.handle = args[0].handle
		else:
			self.handle = _hi.Slice_bool_CTor()
			if len(args) > 0:
				if not isinstance(args[0], collections.Iterable):
					raise TypeError('Slice_bool.__init__ takes a sequence as argument')
				for elt in args[0]:
					self.append(elt)
	def __str__(self):
		s = 'go.Slice_bool len: ' + str(len(self)) + ' handle: ' + str(self.handle) + ' ['
		if len(self) < 120:
			s += ', '.join(map(str, self)) + ']'
		return s
	def __repr__(self):
		return 'go.Slice_bool([' + ', '.join(map(str, self)) + '])'
	def __len__(self):
		return _hi.Slice_bool_len(self.handle)
	def __getitem__(self, key):
		if isinstance(key, slice):
			return [self[ii] for ii in range(*key.indices(len(self)))]
		elif isinstance(key, int):
			if key < 0:
				key += len(self)
			if key < 0 or key >= len(self):
				raise IndexError('slice index out of range')
			return _hi.Slice_bool_elem(self.handle, key)
		else:
			raise TypeError('slice index invalid type')
	def __setitem__(self, idx, value):
		if idx < 0:
			idx += len(self)
		if idx < len(self):
			_hi.Slice_bool_set(self.handle, idx, value)
			return
		raise IndexError('slice index out of range')
	def __iadd__(self, value):
		if not isinstance(value, collections.Iterable):
			raise TypeError('Slice_bool.__iadd__ takes a sequence as argument')
		for elt in value:
			self.append(elt)
		return self
	def __iter__(self):
		self.index = 0
		return self
	def __next__(self):
		if self.index < len(self):
			rv = _hi.Slice_bool_elem(self.handle, self.index)
			self.index = self.index + 1
			return rv
		raise StopIteration
	def append(self, value):
		_hi.Slice_bool_append(self.handle, value)
	def copy(self, src):
		""" copy emulates the go copy function, copying elements into this list from source list, up to min of size of each list """
		mx = min(len(self), len(src))
		for i in range(mx):
			self[i] = src[i]

# Python type for slice []byte
class Slice_byte(GoClass):
	""""""
	def __init__(self, *args, **kwargs):
		"""
		handle=A Go-side object is always initialized with an explicit handle=arg
		otherwise parameter is a python list that we copy from
		"""
		self.index = 0
		if len(kwargs) == 1 and 'handle' in kwargs:
			self.handle = kwargs['handle']
		elif len(args) == 1 and isinstance(args[0], GoClass):
			self.handle = args[0].handle
		else:
			self.handle = _hi.Slice_byte_CTor()
			if len(args) > 0:
				if not isinstance(args[0], collections.Iterable):
					raise TypeError('Slice_byte.__init__ takes a sequence as argument')
				for elt in args[0]:
					self.append(elt)
	def __str__(self):
		s = 'go.Slice_byte len: ' + str(len(self)) + ' handle: ' + str(self.handle) + ' ['
		if len(self) < 120:
			s += ', '.join(map(str, self)) + ']'
		return s
	def __repr__(self):
		return 'go.Slice_byte([' + ', '.join(map(str, self)) + '])'
	def __len__(self):
		return _hi.Slice_byte_len(self.handle)
	def __getitem__(self, key):
		if isinstance(key, slice):
			return [self[ii] for ii in range(*key.indices(len(self)))]
		elif isinstance(key, int):
			if key < 0:
				key += len(self)
			if key < 0 or key >= len(self):
				raise IndexError('slice index out of range')
			return _hi.Slice_byte_elem(self.handle, key)
		else:
			raise TypeError('slice index invalid type')
	def __setitem__(self, idx, value):
		if idx < 0:
			idx += len(self)
		if idx < len(self):
			_hi.Slice_byte_set(self.handle, idx, value)
			return
		raise IndexError('slice index out of range')
	def __iadd__(self, value):
		if not isinstance(value, collections.Iterable):
			raise TypeError('Slice_byte.__iadd__ takes a sequence as argument')
		for elt in value:
			self.append(elt)
		return self
	def __iter__(self):
		self.index = 0
		return self
	def __next__(self):
		if self.index < len(self):
			rv = _hi.Slice_byte_elem(self.handle, self.index)
			self.index = self.index + 1
			return rv
		raise StopIteration
	def append(self, value):
		_hi.Slice_byte_append(self.handle, value)
	def copy(self, src):
		""" copy emulates the go copy function, copying elements into this list from source list, up to min of size of each list """
		mx = min(len(self), len(src))
		for i in range(mx):
			self[i] = src[i]

# Python type for slice []float32
class Slice_float32(GoClass):
	""""""
	def __init__(self, *args, **kwargs):
		"""
		handle=A Go-side object is always initialized with an explicit handle=arg
		otherwise parameter is a python list that we copy from
		"""
		self.index = 0
		if len(kwargs) == 1 and 'handle' in kwargs:
			self.handle = kwargs['handle']
		elif len(args) == 1 and isinstance(args[0], GoClass):
			self.handle = args[0].handle
		else:
			self.handle = _hi.Slice_float32_CTor()
			if len(args) > 0:
				if not isinstance(args[0], collections.Iterable):
					raise TypeError('Slice_float32.__init__ takes a sequence as argument')
				for elt in args[0]:
					self.append(elt)
	def __str__(self):
		s = 'go.Slice_float32 len: ' + str(len(self)) + ' handle: ' + str(self.handle) + ' ['
		if len(self) < 120:
			s += ', '.join(map(str, self)) + ']'
		return s
	def __repr__(self):
		return 'go.Slice_float32([' + ', '.join(map(str, self)) + '])'
	def __len__(self):
		return _hi.Slice_float32_len(self.handle)
	def __getitem__(self, key):
		if isinstance(key, slice):
			return [self[ii] for ii in range(*key.indices(len(self)))]
		elif isinstance(key, int):
			if key < 0:
				key += len(self)
			if key < 0 or key >= len(self):
				raise IndexError('slice index out of range')
			return _hi.Slice_float32_elem(self.handle, key)
		else:
			raise TypeError('slice index invalid type')
	def __setitem__(self, idx, value):
		if idx < 0:
			idx += len(self)
		if idx < len(self):
			_hi.Slice_float32_set(self.handle, idx, value)
			return
		raise IndexError('slice index out of range')
	def __iadd__(self, value):
		if not isinstance(value, collections.Iterable):
			raise TypeError('Slice_float32.__iadd__ takes a sequence as argument')
		for elt in value:
			self.append(elt)
		return self
	def __iter__(self):
		self.index = 0
		return self
	def __next__(self):
		if self.index < len(self):
			rv = _hi.Slice_float32_elem(self.handle, self.index)
			self.index = self.index + 1
			return rv
		raise StopIteration
	def append(self, value):
		_hi.Slice_float32_append(self.handle, value)
	def copy(self, src):
		""" copy emulates the go copy function, copying elements into this list from source list, up to min of size of each list """
		mx = min(len(self), len(src))
		for i in range(mx):
			self[i] = src[i]

# Python type for slice []float64
class Slice_float64(GoClass):
	""""""
	def __init__(self, *args, **kwargs):
		"""
		handle=A Go-side object is always initialized with an explicit handle=arg
		otherwise parameter is a python list that we copy from
		"""
		self.index = 0
		if len(kwargs) == 1 and 'handle' in kwargs:
			self.handle = kwargs['handle']
		elif len(args) == 1 and isinstance(args[0], GoClass):
			self.handle = args[0].handle
		else:
			self.handle = _hi.Slice_float64_CTor()
			if len(args) > 0:
				if not isinstance(args[0], collections.Iterable):
					raise TypeError('Slice_float64.__init__ takes a sequence as argument')
				for elt in args[0]:
					self.append(elt)
	def __str__(self):
		s = 'go.Slice_float64 len: ' + str(len(self)) + ' handle: ' + str(self.handle) + ' ['
		if len(self) < 120:
			s += ', '.join(map(str, self)) + ']'
		return s
	def __repr__(self):
		return 'go.Slice_float64([' + ', '.join(map(str, self)) + '])'
	def __len__(self):
		return _hi.Slice_float64_len(self.handle)
	def __getitem__(self, key):
		if isinstance(key, slice):
			return [self[ii] for ii in range(*key.indices(len(self)))]
		elif isinstance(key, int):
			if key < 0:
				key += len(self)
			if key < 0 or key >= len(self):
				raise IndexError('slice index out of range')
			return _hi.Slice_float64_elem(self.handle, key)
		else:
			raise TypeError('slice index invalid type')
	def __setitem__(self, idx, value):
		if idx < 0:
			idx += len(self)
		if idx < len(self):
			_hi.Slice_float64_set(self.handle, idx, value)
			return
		raise IndexError('slice index out of range')
	def __iadd__(self, value):
		if not isinstance(value, collections.Iterable):
			raise TypeError('Slice_float64.__iadd__ takes a sequence as argument')
		for elt in value:
			self.append(elt)
		return self
	def __iter__(self):
		self.index = 0
		return self
	def __next__(self):
		if self.index < len(self):
			rv = _hi.Slice_float64_elem(self.handle, self.index)
			self.index = self.index + 1
			return rv
		raise StopIteration
	def append(self, value):
		_hi.Slice_float64_append(self.handle, value)
	def copy(self, src):
		""" copy emulates the go copy function, copying elements into this list from source list, up to min of size of each list """
		mx = min(len(self), len(src))
		for i in range(mx):
			self[i] = src[i]

# Python type for slice []int
class Slice_int(GoClass):
	""""""
	def __init__(self, *args, **kwargs):
		"""
		handle=A Go-side object is always initialized with an explicit handle=arg
		otherwise parameter is a python list that we copy from
		"""
		self.index = 0
		if len(kwargs) == 1 and 'handle' in kwargs:
			self.handle = kwargs['handle']
		elif len(args) == 1 and isinstance(args[0], GoClass):
			self.handle = args[0].handle
		else:
			self.handle = _hi.Slice_int_CTor()
			if len(args) > 0:
				if not isinstance(args[0], collections.Iterable):
					raise TypeError('Slice_int.__init__ takes a sequence as argument')
				for elt in args[0]:
					self.append(elt)
	def __str__(self):
		s = 'go.Slice_int len: ' + str(len(self)) + ' handle: ' + str(self.handle) + ' ['
		if len(self) < 120:
			s += ', '.join(map(str, self)) + ']'
		return s
	def __repr__(self):
		return 'go.Slice_int([' + ', '.join(map(str, self)) + '])'
	def __len__(self):
		return _hi.Slice_int_len(self.handle)
	def __getitem__(self, key):
		if isinstance(key, slice):
			return [self[ii] for ii in range(*key.indices(len(self)))]
		elif isinstance(key, int):
			if key < 0:
				key += len(self)
			if key < 0 or key >= len(self):
				raise IndexError('slice index out of range')
			return _hi.Slice_int_elem(self.handle, key)
		else:
			raise TypeError('slice index invalid type')
	def __setitem__(self, idx, value):
		if idx < 0:
			idx += len(self)
		if idx < len(self):
			_hi.Slice_int_set(self.handle, idx, value)
			return
		raise IndexError('slice index out of range')
	def __iadd__(self, value):
		if not isinstance(value, collections.Iterable):
			raise TypeError('Slice_int.__iadd__ takes a sequence as argument')
		for elt in value:
			self.append(elt)
		return self
	def __iter__(self):
		self.index = 0
		return self
	def __next__(self):
		if self.index < len(self):
			rv = _hi.Slice_int_elem(self.handle, self.index)
			self.index = self.index + 1
			return rv
		raise StopIteration
	def append(self, value):
		_hi.Slice_int_append(self.handle, value)
	def copy(self, src):
		""" copy emulates the go copy function, copying elements into this list from source list, up to min of size of each list """
		mx = min(len(self), len(src))
		for i in range(mx):
			self[i] = src[i]

# Python type for slice []int16
class Slice_int16(GoClass):
	""""""
	def __init__(self, *args, **kwargs):
		"""
		handle=A Go-side object is always initialized with an explicit handle=arg
		otherwise parameter is a python list that we copy from
		"""
		self.index = 0
		if len(kwargs) == 1 and 'handle' in kwargs:
			self.handle = kwargs['handle']
		elif len(args) == 1 and isinstance(args[0], GoClass):
			self.handle = args[0].handle
		else:
			self.handle = _hi.Slice_int16_CTor()
			if len(args) > 0:
				if not isinstance(args[0], collections.Iterable):
					raise TypeError('Slice_int16.__init__ takes a sequence as argument')
				for elt in args[0]:
					self.append(elt)
	def __str__(self):
		s = 'go.Slice_int16 len: ' + str(len(self)) + ' handle: ' + str(self.handle) + ' ['
		if len(self) < 120:
			s += ', '.join(map(str, self)) + ']'
		return s
	def __repr__(self):
		return 'go.Slice_int16([' + ', '.join(map(str, self)) + '])'
	def __len__(self):
		return _hi.Slice_int16_len(self.handle)
	def __getitem__(self, key):
		if isinstance(key, slice):
			return [self[ii] for ii in range(*key.indices(len(self)))]
		elif isinstance(key, int):
			if key < 0:
				key += len(self)
			if key < 0 or key >= len(self):
				raise IndexError('slice index out of range')
			return _hi.Slice_int16_elem(self.handle, key)
		else:
			raise TypeError('slice index invalid type')
	def __setitem__(self, idx, value):
		if idx < 0:
			idx += len(self)
		if idx < len(self):
			_hi.Slice_int16_set(self.handle, idx, value)
			return
		raise IndexError('slice index out of range')
	def __iadd__(self, value):
		if not isinstance(value, collections.Iterable):
			raise TypeError('Slice_int16.__iadd__ takes a sequence as argument')
		for elt in value:
			self.append(elt)
		return self
	def __iter__(self):
		self.index = 0
		return self
	def __next__(self):
		if self.index < len(self):
			rv = _hi.Slice_int16_elem(self.handle, self.index)
			self.index = self.index + 1
			return rv
		raise StopIteration
	def append(self, value):
		_hi.Slice_int16_append(self.handle, value)
	def copy(self, src):
		""" copy emulates the go copy function, copying elements into this list from source list, up to min of size of each list """
		mx = min(len(self), len(src))
		for i in range(mx):
			self[i] = src[i]

# Python type for slice []int32
class Slice_int32(GoClass):
	""""""
	def __init__(self, *args, **kwargs):
		"""
		handle=A Go-side object is always initialized with an explicit handle=arg
		otherwise parameter is a python list that we copy from
		"""
		self.index = 0
		if len(kwargs) == 1 and 'handle' in kwargs:
			self.handle = kwargs['handle']
		elif len(args) == 1 and isinstance(args[0], GoClass):
			self.handle = args[0].handle
		else:
			self.handle = _hi.Slice_int32_CTor()
			if len(args) > 0:
				if not isinstance(args[0], collections.Iterable):
					raise TypeError('Slice_int32.__init__ takes a sequence as argument')
				for elt in args[0]:
					self.append(elt)
	def __str__(self):
		s = 'go.Slice_int32 len: ' + str(len(self)) + ' handle: ' + str(self.handle) + ' ['
		if len(self) < 120:
			s += ', '.join(map(str, self)) + ']'
		return s
	def __repr__(self):
		return 'go.Slice_int32([' + ', '.join(map(str, self)) + '])'
	def __len__(self):
		return _hi.Slice_int32_len(self.handle)
	def __getitem__(self, key):
		if isinstance(key, slice):
			return [self[ii] for ii in range(*key.indices(len(self)))]
		elif isinstance(key, int):
			if key < 0:
				key += len(self)
			if key < 0 or key >= len(self):
				raise IndexError('slice index out of range')
			return _hi.Slice_int32_elem(self.handle, key)
		else:
			raise TypeError('slice index invalid type')
	def __setitem__(self, idx, value):
		if idx < 0:
			idx += len(self)
		if idx < len(self):
			_hi.Slice_int32_set(self.handle, idx, value)
			return
		raise IndexError('slice index out of range')
	def __iadd__(self, value):
		if not isinstance(value, collections.Iterable):
			raise TypeError('Slice_int32.__iadd__ takes a sequence as argument')
		for elt in value:
			self.append(elt)
		return self
	def __iter__(self):
		self.index = 0
		return self
	def __next__(self):
		if self.index < len(self):
			rv = _hi.Slice_int32_elem(self.handle, self.index)
			self.index = self.index + 1
			return rv
		raise StopIteration
	def append(self, value):
		_hi.Slice_int32_append(self.handle, value)
	def copy(self, src):
		""" copy emulates the go copy function, copying elements into this list from source list, up to min of size of each list """
		mx = min(len(self), len(src))
		for i in range(mx):
			self[i] = src[i]

# Python type for slice []int64
class Slice_int64(GoClass):
	""""""
	def __init__(self, *args, **kwargs):
		"""
		handle=A Go-side object is always initialized with an explicit handle=arg
		otherwise parameter is a python list that we copy from
		"""
		self.index = 0
		if len(kwargs) == 1 and 'handle' in kwargs:
			self.handle = kwargs['handle']
		elif len(args) == 1 and isinstance(args[0], GoClass):
			self.handle = args[0].handle
		else:
			self.handle = _hi.Slice_int64_CTor()
			if len(args) > 0:
				if not isinstance(args[0], collections.Iterable):
					raise TypeError('Slice_int64.__init__ takes a sequence as argument')
				for elt in args[0]:
					self.append(elt)
	def __str__(self):
		s = 'go.Slice_int64 len: ' + str(len(self)) + ' handle: ' + str(self.handle) + ' ['
		if len(self) < 120:
			s += ', '.join(map(str, self)) + ']'
		return s
	def __repr__(self):
		return 'go.Slice_int64([' + ', '.join(map(str, self)) + '])'
	def __len__(self):
		return _hi.Slice_int64_len(self.handle)
	def __getitem__(self, key):
		if isinstance(key, slice):
			return [self[ii] for ii in range(*key.indices(len(self)))]
		elif isinstance(key, int):
			if key < 0:
				key += len(self)
			if key < 0 or key >= len(self):
				raise IndexError('slice index out of range')
			return _hi.Slice_int64_elem(self.handle, key)
		else:
			raise TypeError('slice index invalid type')
	def __setitem__(self, idx, value):
		if idx < 0:
			idx += len(self)
		if idx < len(self):
			_hi.Slice_int64_set(self.handle, idx, value)
			return
		raise IndexError('slice index out of range')
	def __iadd__(self, value):
		if not isinstance(value, collections.Iterable):
			raise TypeError('Slice_int64.__iadd__ takes a sequence as argument')
		for elt in value:
			self.append(elt)
		return self
	def __iter__(self):
		self.index = 0
		return self
	def __next__(self):
		if self.index < len(self):
			rv = _hi.Slice_int64_elem(self.handle, self.index)
			self.index = self.index + 1
			return rv
		raise StopIteration
	def append(self, value):
		_hi.Slice_int64_append(self.handle, value)
	def copy(self, src):
		""" copy emulates the go copy function, copying elements into this list from source list, up to min of size of each list """
		mx = min(len(self), len(src))
		for i in range(mx):
			self[i] = src[i]

# Python type for slice []int8
class Slice_int8(GoClass):
	""""""
	def __init__(self, *args, **kwargs):
		"""
		handle=A Go-side object is always initialized with an explicit handle=arg
		otherwise parameter is a python list that we copy from
		"""
		self.index = 0
		if len(kwargs) == 1 and 'handle' in kwargs:
			self.handle = kwargs['handle']
		elif len(args) == 1 and isinstance(args[0], GoClass):
			self.handle = args[0].handle
		else:
			self.handle = _hi.Slice_int8_CTor()
			if len(args) > 0:
				if not isinstance(args[0], collections.Iterable):
					raise TypeError('Slice_int8.__init__ takes a sequence as argument')
				for elt in args[0]:
					self.append(elt)
	def __str__(self):
		s = 'go.Slice_int8 len: ' + str(len(self)) + ' handle: ' + str(self.handle) + ' ['
		if len(self) < 120:
			s += ', '.join(map(str, self)) + ']'
		return s
	def __repr__(self):
		return 'go.Slice_int8([' + ', '.join(map(str, self)) + '])'
	def __len__(self):
		return _hi.Slice_int8_len(self.handle)
	def __getitem__(self, key):
		if isinstance(key, slice):
			return [self[ii] for ii in range(*key.indices(len(self)))]
		elif isinstance(key, int):
			if key < 0:
				key += len(self)
			if key < 0 or key >= len(self):
				raise IndexError('slice index out of range')
			return _hi.Slice_int8_elem(self.handle, key)
		else:
			raise TypeError('slice index invalid type')
	def __setitem__(self, idx, value):
		if idx < 0:
			idx += len(self)
		if idx < len(self):
			_hi.Slice_int8_set(self.handle, idx, value)
			return
		raise IndexError('slice index out of range')
	def __iadd__(self, value):
		if not isinstance(value, collections.Iterable):
			raise TypeError('Slice_int8.__iadd__ takes a sequence as argument')
		for elt in value:
			self.append(elt)
		return self
	def __iter__(self):
		self.index = 0
		return self
	def __next__(self):
		if self.index < len(self):
			rv = _hi.Slice_int8_elem(self.handle, self.index)
			self.index = self.index + 1
			return rv
		raise StopIteration
	def append(self, value):
		_hi.Slice_int8_append(self.handle, value)
	def copy(self, src):
		""" copy emulates the go copy function, copying elements into this list from source list, up to min of size of each list """
		mx = min(len(self), len(src))
		for i in range(mx):
			self[i] = src[i]

# Python type for slice []rune
class Slice_rune(GoClass):
	""""""
	def __init__(self, *args, **kwargs):
		"""
		handle=A Go-side object is always initialized with an explicit handle=arg
		otherwise parameter is a python list that we copy from
		"""
		self.index = 0
		if len(kwargs) == 1 and 'handle' in kwargs:
			self.handle = kwargs['handle']
		elif len(args) == 1 and isinstance(args[0], GoClass):
			self.handle = args[0].handle
		else:
			self.handle = _hi.Slice_rune_CTor()
			if len(args) > 0:
				if not isinstance(args[0], collections.Iterable):
					raise TypeError('Slice_rune.__init__ takes a sequence as argument')
				for elt in args[0]:
					self.append(elt)
	def __str__(self):
		s = 'go.Slice_rune len: ' + str(len(self)) + ' handle: ' + str(self.handle) + ' ['
		if len(self) < 120:
			s += ', '.join(map(str, self)) + ']'
		return s
	def __repr__(self):
		return 'go.Slice_rune([' + ', '.join(map(str, self)) + '])'
	def __len__(self):
		return _hi.Slice_rune_len(self.handle)
	def __getitem__(self, key):
		if isinstance(key, slice):
			return [self[ii] for ii in range(*key.indices(len(self)))]
		elif isinstance(key, int):
			if key < 0:
				key += len(self)
			if key < 0 or key >= len(self):
				raise IndexError('slice index out of range')
			return _hi.Slice_rune_elem(self.handle, key)
		else:
			raise TypeError('slice index invalid type')
	def __setitem__(self, idx, value):
		if idx < 0:
			idx += len(self)
		if idx < len(self):
			_hi.Slice_rune_set(self.handle, idx, value)
			return
		raise IndexError('slice index out of range')
	def __iadd__(self, value):
		if not isinstance(value, collections.Iterable):
			raise TypeError('Slice_rune.__iadd__ takes a sequence as argument')
		for elt in value:
			self.append(elt)
		return self
	def __iter__(self):
		self.index = 0
		return self
	def __next__(self):
		if self.index < len(self):
			rv = _hi.Slice_rune_elem(self.handle, self.index)
			self.index = self.index + 1
			return rv
		raise StopIteration
	def append(self, value):
		_hi.Slice_rune_append(self.handle, value)
	def copy(self, src):
		""" copy emulates the go copy function, copying elements into this list from source list, up to min of size of each list """
		mx = min(len(self), len(src))
		for i in range(mx):
			self[i] = src[i]

# Python type for slice []string
class Slice_string(GoClass):
	""""""
	def __init__(self, *args, **kwargs):
		"""
		handle=A Go-side object is always initialized with an explicit handle=arg
		otherwise parameter is a python list that we copy from
		"""
		self.index = 0
		if len(kwargs) == 1 and 'handle' in kwargs:
			self.handle = kwargs['handle']
		elif len(args) == 1 and isinstance(args[0], GoClass):
			self.handle = args[0].handle
		else:
			self.handle = _hi.Slice_string_CTor()
			if len(args) > 0:
				if not isinstance(args[0], collections.Iterable):
					raise TypeError('Slice_string.__init__ takes a sequence as argument')
				for elt in args[0]:
					self.append(elt)
	def __str__(self):
		s = 'go.Slice_string len: ' + str(len(self)) + ' handle: ' + str(self.handle) + ' ['
		if len(self) < 120:
			s += ', '.join(map(str, self)) + ']'
		return s
	def __repr__(self):
		return 'go.Slice_string([' + ', '.join(map(str, self)) + '])'
	def __len__(self):
		return _hi.Slice_string_len(self.handle)
	def __getitem__(self, key):
		if isinstance(key, slice):
			return [self[ii] for ii in range(*key.indices(len(self)))]
		elif isinstance(key, int):
			if key < 0:
				key += len(self)
			if key < 0 or key >= len(self):
				raise IndexError('slice index out of range')
			return _hi.Slice_string_elem(self.handle, key)
		else:
			raise TypeError('slice index invalid type')
	def __setitem__(self, idx, value):
		if idx < 0:
			idx += len(self)
		if idx < len(self):
			_hi.Slice_string_set(self.handle, idx, value)
			return
		raise IndexError('slice index out of range')
	def __iadd__(self, value):
		if not isinstance(value, collections.Iterable):
			raise TypeError('Slice_string.__iadd__ takes a sequence as argument')
		for elt in value:
			self.append(elt)
		return self
	def __iter__(self):
		self.index = 0
		return self
	def __next__(self):
		if self.index < len(self):
			rv = _hi.Slice_string_elem(self.handle, self.index)
			self.index = self.index + 1
			return rv
		raise StopIteration
	def append(self, value):
		_hi.Slice_string_append(self.handle, value)
	def copy(self, src):
		""" copy emulates the go copy function, copying elements into this list from source list, up to min of size of each list """
		mx = min(len(self), len(src))
		for i in range(mx):
			self[i] = src[i]

# Python type for slice []uint
class Slice_uint(GoClass):
	""""""
	def __init__(self, *args, **kwargs):
		"""
		handle=A Go-side object is always initialized with an explicit handle=arg
		otherwise parameter is a python list that we copy from
		"""
		self.index = 0
		if len(kwargs) == 1 and 'handle' in kwargs:
			self.handle = kwargs['handle']
		elif len(args) == 1 and isinstance(args[0], GoClass):
			self.handle = args[0].handle
		else:
			self.handle = _hi.Slice_uint_CTor()
			if len(args) > 0:
				if not isinstance(args[0], collections.Iterable):
					raise TypeError('Slice_uint.__init__ takes a sequence as argument')
				for elt in args[0]:
					self.append(elt)
	def __str__(self):
		s = 'go.Slice_uint len: ' + str(len(self)) + ' handle: ' + str(self.handle) + ' ['
		if len(self) < 120:
			s += ', '.join(map(str, self)) + ']'
		return s
	def __repr__(self):
		return 'go.Slice_uint([' + ', '.join(map(str, self)) + '])'
	def __len__(self):
		return _hi.Slice_uint_len(self.handle)
	def __getitem__(self, key):
		if isinstance(key, slice):
			return [self[ii] for ii in range(*key.indices(len(self)))]
		elif isinstance(key, int):
			if key < 0:
				key += len(self)
			if key < 0 or key >= len(self):
				raise IndexError('slice index out of range')
			return _hi.Slice_uint_elem(self.handle, key)
		else:
			raise TypeError('slice index invalid type')
	def __setitem__(self, idx, value):
		if idx < 0:
			idx += len(self)
		if idx < len(self):
			_hi.Slice_uint_set(self.handle, idx, value)
			return
		raise IndexError('slice index out of range')
	def __iadd__(self, value):
		if not isinstance(value, collections.Iterable):
			raise TypeError('Slice_uint.__iadd__ takes a sequence as argument')
		for elt in value:
			self.append(elt)
		return self
	def __iter__(self):
		self.index = 0
		return self
	def __next__(self):
		if self.index < len(self):
			rv = _hi.Slice_uint_elem(self.handle, self.index)
			self.index = self.index + 1
			return rv
		raise StopIteration
	def append(self, value):
		_hi.Slice_uint_append(self.handle, value)
	def copy(self, src):
		""" copy emulates the go copy function, copying elements into this list from source list, up to min of size of each list """
		mx = min(len(self), len(src))
		for i in range(mx):
			self[i] = src[i]

# Python type for slice []uint16
class Slice_uint16(GoClass):
	""""""
	def __init__(self, *args, **kwargs):
		"""
		handle=A Go-side object is always initialized with an explicit handle=arg
		otherwise parameter is a python list that we copy from
		"""
		self.index = 0
		if len(kwargs) == 1 and 'handle' in kwargs:
			self.handle = kwargs['handle']
		elif len(args) == 1 and isinstance(args[0], GoClass):
			self.handle = args[0].handle
		else:
			self.handle = _hi.Slice_uint16_CTor()
			if len(args) > 0:
				if not isinstance(args[0], collections.Iterable):
					raise TypeError('Slice_uint16.__init__ takes a sequence as argument')
				for elt in args[0]:
					self.append(elt)
	def __str__(self):
		s = 'go.Slice_uint16 len: ' + str(len(self)) + ' handle: ' + str(self.handle) + ' ['
		if len(self) < 120:
			s += ', '.join(map(str, self)) + ']'
		return s
	def __repr__(self):
		return 'go.Slice_uint16([' + ', '.join(map(str, self)) + '])'
	def __len__(self):
		return _hi.Slice_uint16_len(self.handle)
	def __getitem__(self, key):
		if isinstance(key, slice):
			return [self[ii] for ii in range(*key.indices(len(self)))]
		elif isinstance(key, int):
			if key < 0:
				key += len(self)
			if key < 0 or key >= len(self):
				raise IndexError('slice index out of range')
			return _hi.Slice_uint16_elem(self.handle, key)
		else:
			raise TypeError('slice index invalid type')
	def __setitem__(self, idx, value):
		if idx < 0:
			idx += len(self)
		if idx < len(self):
			_hi.Slice_uint16_set(self.handle, idx, value)
			return
		raise IndexError('slice index out of range')
	def __iadd__(self, value):
		if not isinstance(value, collections.Iterable):
			raise TypeError('Slice_uint16.__iadd__ takes a sequence as argument')
		for elt in value:
			self.append(elt)
		return self
	def __iter__(self):
		self.index = 0
		return self
	def __next__(self):
		if self.index < len(self):
			rv = _hi.Slice_uint16_elem(self.handle, self.index)
			self.index = self.index + 1
			return rv
		raise StopIteration
	def append(self, value):
		_hi.Slice_uint16_append(self.handle, value)
	def copy(self, src):
		""" copy emulates the go copy function, copying elements into this list from source list, up to min of size of each list """
		mx = min(len(self), len(src))
		for i in range(mx):
			self[i] = src[i]

# Python type for slice []uint32
class Slice_uint32(GoClass):
	""""""
	def __init__(self, *args, **kwargs):
		"""
		handle=A Go-side object is always initialized with an explicit handle=arg
		otherwise parameter is a python list that we copy from
		"""
		self.index = 0
		if len(kwargs) == 1 and 'handle' in kwargs:
			self.handle = kwargs['handle']
		elif len(args) == 1 and isinstance(args[0], GoClass):
			self.handle = args[0].handle
		else:
			self.handle = _hi.Slice_uint32_CTor()
			if len(args) > 0:
				if not isinstance(args[0], collections.Iterable):
					raise TypeError('Slice_uint32.__init__ takes a sequence as argument')
				for elt in args[0]:
					self.append(elt)
	def __str__(self):
		s = 'go.Slice_uint32 len: ' + str(len(self)) + ' handle: ' + str(self.handle) + ' ['
		if len(self) < 120:
			s += ', '.join(map(str, self)) + ']'
		return s
	def __repr__(self):
		return 'go.Slice_uint32([' + ', '.join(map(str, self)) + '])'
	def __len__(self):
		return _hi.Slice_uint32_len(self.handle)
	def __getitem__(self, key):
		if isinstance(key, slice):
			return [self[ii] for ii in range(*key.indices(len(self)))]
		elif isinstance(key, int):
			if key < 0:
				key += len(self)
			if key < 0 or key >= len(self):
				raise IndexError('slice index out of range')
			return _hi.Slice_uint32_elem(self.handle, key)
		else:
			raise TypeError('slice index invalid type')
	def __setitem__(self, idx, value):
		if idx < 0:
			idx += len(self)
		if idx < len(self):
			_hi.Slice_uint32_set(self.handle, idx, value)
			return
		raise IndexError('slice index out of range')
	def __iadd__(self, value):
		if not isinstance(value, collections.Iterable):
			raise TypeError('Slice_uint32.__iadd__ takes a sequence as argument')
		for elt in value:
			self.append(elt)
		return self
	def __iter__(self):
		self.index = 0
		return self
	def __next__(self):
		if self.index < len(self):
			rv = _hi.Slice_uint32_elem(self.handle, self.index)
			self.index = self.index + 1
			return rv
		raise StopIteration
	def append(self, value):
		_hi.Slice_uint32_append(self.handle, value)
	def copy(self, src):
		""" copy emulates the go copy function, copying elements into this list from source list, up to min of size of each list """
		mx = min(len(self), len(src))
		for i in range(mx):
			self[i] = src[i]

# Python type for slice []uint64
class Slice_uint64(GoClass):
	""""""
	def __init__(self, *args, **kwargs):
		"""
		handle=A Go-side object is always initialized with an explicit handle=arg
		otherwise parameter is a python list that we copy from
		"""
		self.index = 0
		if len(kwargs) == 1 and 'handle' in kwargs:
			self.handle = kwargs['handle']
		elif len(args) == 1 and isinstance(args[0], GoClass):
			self.handle = args[0].handle
		else:
			self.handle = _hi.Slice_uint64_CTor()
			if len(args) > 0:
				if not isinstance(args[0], collections.Iterable):
					raise TypeError('Slice_uint64.__init__ takes a sequence as argument')
				for elt in args[0]:
					self.append(elt)
	def __str__(self):
		s = 'go.Slice_uint64 len: ' + str(len(self)) + ' handle: ' + str(self.handle) + ' ['
		if len(self) < 120:
			s += ', '.join(map(str, self)) + ']'
		return s
	def __repr__(self):
		return 'go.Slice_uint64([' + ', '.join(map(str, self)) + '])'
	def __len__(self):
		return _hi.Slice_uint64_len(self.handle)
	def __getitem__(self, key):
		if isinstance(key, slice):
			return [self[ii] for ii in range(*key.indices(len(self)))]
		elif isinstance(key, int):
			if key < 0:
				key += len(self)
			if key < 0 or key >= len(self):
				raise IndexError('slice index out of range')
			return _hi.Slice_uint64_elem(self.handle, key)
		else:
			raise TypeError('slice index invalid type')
	def __setitem__(self, idx, value):
		if idx < 0:
			idx += len(self)
		if idx < len(self):
			_hi.Slice_uint64_set(self.handle, idx, value)
			return
		raise IndexError('slice index out of range')
	def __iadd__(self, value):
		if not isinstance(value, collections.Iterable):
			raise TypeError('Slice_uint64.__iadd__ takes a sequence as argument')
		for elt in value:
			self.append(elt)
		return self
	def __iter__(self):
		self.index = 0
		return self
	def __next__(self):
		if self.index < len(self):
			rv = _hi.Slice_uint64_elem(self.handle, self.index)
			self.index = self.index + 1
			return rv
		raise StopIteration
	def append(self, value):
		_hi.Slice_uint64_append(self.handle, value)
	def copy(self, src):
		""" copy emulates the go copy function, copying elements into this list from source list, up to min of size of each list """
		mx = min(len(self), len(src))
		for i in range(mx):
			self[i] = src[i]

# Python type for slice []uint8
class Slice_uint8(GoClass):
	""""""
	def __init__(self, *args, **kwargs):
		"""
		handle=A Go-side object is always initialized with an explicit handle=arg
		otherwise parameter is a python list that we copy from
		"""
		self.index = 0
		if len(kwargs) == 1 and 'handle' in kwargs:
			self.handle = kwargs['handle']
		elif len(args) == 1 and isinstance(args[0], GoClass):
			self.handle = args[0].handle
		else:
			self.handle = _hi.Slice_uint8_CTor()
			if len(args) > 0:
				if not isinstance(args[0], collections.Iterable):
					raise TypeError('Slice_uint8.__init__ takes a sequence as argument')
				for elt in args[0]:
					self.append(elt)
	def __str__(self):
		s = 'go.Slice_uint8 len: ' + str(len(self)) + ' handle: ' + str(self.handle) + ' ['
		if len(self) < 120:
			s += ', '.join(map(str, self)) + ']'
		return s
	def __repr__(self):
		return 'go.Slice_uint8([' + ', '.join(map(str, self)) + '])'
	def __len__(self):
		return _hi.Slice_uint8_len(self.handle)
	def __getitem__(self, key):
		if isinstance(key, slice):
			return [self[ii] for ii in range(*key.indices(len(self)))]
		elif isinstance(key, int):
			if key < 0:
				key += len(self)
			if key < 0 or key >= len(self):
				raise IndexError('slice index out of range')
			return _hi.Slice_uint8_elem(self.handle, key)
		else:
			raise TypeError('slice index invalid type')
	def __setitem__(self, idx, value):
		if idx < 0:
			idx += len(self)
		if idx < len(self):
			_hi.Slice_uint8_set(self.handle, idx, value)
			return
		raise IndexError('slice index out of range')
	def __iadd__(self, value):
		if not isinstance(value, collections.Iterable):
			raise TypeError('Slice_uint8.__iadd__ takes a sequence as argument')
		for elt in value:
			self.append(elt)
		return self
	def __iter__(self):
		self.index = 0
		return self
	def __next__(self):
		if self.index < len(self):
			rv = _hi.Slice_uint8_elem(self.handle, self.index)
			self.index = self.index + 1
			return rv
		raise StopIteration
	def append(self, value):
		_hi.Slice_uint8_append(self.handle, value)
	def copy(self, src):
		""" copy emulates the go copy function, copying elements into this list from source list, up to min of size of each list """
		mx = min(len(self), len(src))
		for i in range(mx):
			self[i] = src[i]

# ---- External Types Outside of Targeted Packages ---

# Python type for *structs.S2
class Ptr_structs_S2(GoClass):
	""""""
	def __init__(self, *args, **kwargs):
		"""
		handle=A Go-side object is always initialized with an explicit handle=arg
		"""
		if len(kwargs) == 1 and 'handle' in kwargs:
			self.handle = kwargs['handle']
		elif len(args) == 1 and isinstance(args[0], GoClass):
			self.handle = args[0].handle
		elif len(args) == 1 and isinstance(args[0], int):
			self.handle = args[0]
		else:
			self.handle = 0
	

# Python type for structs.S2
class structs_S2(GoClass):
	""""""
	def __init__(self, *args, **kwargs):
		"""
		handle=A Go-side object is always initialized with an explicit handle=arg
		"""
		if len(kwargs) == 1 and 'handle' in kwargs:
			self.handle = kwargs['handle']
		elif len(args) == 1 and isinstance(args[0], GoClass):
			self.handle = args[0].handle
		elif len(args) == 1 and isinstance(args[0], int):
			self.handle = args[0]
		else:
			self.handle = 0
	


